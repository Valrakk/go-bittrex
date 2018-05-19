package bittrex

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"github.com/thebotguys/signalr"
)

// OrderUpdate are the changes in the order book
type OrderUpdate struct {
	Quantity decimal.Decimal `json:"Quantity"`
	Rate     decimal.Decimal `json:"Rate"`
	Type     int             `json:"Type"`
}

// InitialFill is used for the first update frame
type InitialFill struct {
	ID        int             `json:"Id"`
	Timestamp jTime           `json:"Timestamp"`
	Quantity  decimal.Decimal `json:"Quantity"`
	Price     decimal.Decimal `json:"Price"`
	Total     decimal.Decimal `json:"Total"`
	FillType  string          `json:"FillType"`
	OrderType string          `json:"OrderType"`
}

// Fill are the executed orders
type Fill struct {
	OrderType string          `json:"OrderType"`
	Quantity  decimal.Decimal `json:"Quantity"`
	Rate      decimal.Decimal `json:"Rate"`
	Timestamp jTime           `json:"Timestamp"`
}

// ExchangeState contains fills and order book updates for a market.
type ExchangeState struct {
	MarketName string        `json:"MarketName"`
	Nounce     int           `json:"Nounce"`
	Buys       []OrderUpdate `json:"Buys"`
	Sells      []OrderUpdate `json:"Sells"`
	Fills      []Fill        `json:"Fills"`
	Initial    bool
}

// InitialExchangeState contains the initial fills and order book updates for a market.
type InitialExchangeState struct {
	MarketName string        `json:"MarketName"`
	Nounce     int           `json:"Nounce"`
	Buys       []OrderUpdate `json:"Buys"`
	Sells      []OrderUpdate `json:"Sells"`
	Fills      []InitialFill `json:"Fills"`
	Initial    bool
}

// doAsyncTimeout runs f in a different goroutine
//	if f returns before timeout elapses, doAsyncTimeout returns the result of f().
//	otherwise it returns "operation timeout" error, and calls tmFunc after f returns.
func doAsyncTimeout(f func() error, tmFunc func(error), timeout time.Duration) error {
	errs := make(chan error)
	go func() {
		err := f()
		select {
		case errs <- err:
		default:
			if tmFunc != nil {
				tmFunc(err)
			}
		}
	}()
	select {
	case err := <-errs:
		return err
	case <-time.After(timeout):
		return errors.New("operation timeout")
	}
}

func sendStateAsync(dataCh chan<- ExchangeState, st ExchangeState) {
	select {
	case dataCh <- st:
	default:
	}
}

func subForMarket(client *signalr.Client, market string) (json.RawMessage, error) {
	_, err := client.CallHub(wsHub, "SubscribeToExchangeDeltas", market)
	if err != nil {
		return json.RawMessage{}, err
	}
	return client.CallHub(wsHub, "QueryExchangeState", market)
}

func parseStates(messages []json.RawMessage, dataCh chan<- ExchangeState, market string) {
	for _, msg := range messages {
		var st ExchangeState
		if err := json.Unmarshal(msg, &st); err != nil {
			continue
		}
		if st.MarketName != market {
			continue
		}
		sendStateAsync(dataCh, st)
	}
}

// SubscribeExchangeUpdate subscribes for updates of the market.
// Updates will be sent to dataCh.
// To stop subscription, send to, or close 'stop'.
func (b *Bittrex) SubscribeExchangeUpdate(market string, dataCh chan<- ExchangeState, stop <-chan bool) error {
	const timeout = 5 * time.Second
	client := signalr.NewWebsocketClient()
	client.OnClientMethod = func(hub string, method string, messages []json.RawMessage) {
		if hub != wsHub || method != "updateExchangeState" {
			return
		}
		parseStates(messages, dataCh, market)
	}
	err := doAsyncTimeout(func() error {
		return client.Connect("https", wsBase, []string{wsHub})
	}, func(err error) {
		if err == nil {
			client.Close()
		}
	}, timeout)
	if err != nil {
		return err
	}
	defer client.Close()
	var msg json.RawMessage
	err = doAsyncTimeout(func() error {
		var err error
		msg, err = subForMarket(client, market)
		return err
	}, nil, timeout)
	if err != nil {
		return err
	}
	var ist InitialExchangeState
	var st ExchangeState
	if err = json.Unmarshal(msg, &ist); err != nil {
		return err
	}

	st.Initial = true
	st.MarketName = market
	st.Buys = make([]OrderUpdate, len(ist.Buys))
	st.Sells = make([]OrderUpdate, len(ist.Sells))
	st.Fills = make([]Fill, len(ist.Fills))

	for i := range ist.Buys {
		st.Buys[i] = ist.Buys[i]
	}
	for i := range ist.Sells {
		st.Sells[i] = ist.Sells[i]
	}
	for i := range ist.Fills {
		st.Fills[i].OrderType = ist.Fills[i].OrderType
		st.Fills[i].Quantity = ist.Fills[i].Quantity
		st.Fills[i].Rate = ist.Fills[i].Price
		st.Fills[i].Timestamp = ist.Fills[i].Timestamp
	}

	sendStateAsync(dataCh, st)
	select {
	case <-stop:
	case <-client.DisconnectedChannel:
	}
	return nil
}
