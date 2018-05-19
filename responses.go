package bittrex

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

// jsonResponse .
type jsonResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

//Address .
type Address struct {
	Currency string `json:"Currency"`
	Address  string `json:"Address"`
}

//Balance .
type Balance struct {
	Currency      string          `json:"Currency"`
	Balance       decimal.Decimal `json:"Balance"`
	Available     decimal.Decimal `json:"Available"`
	Pending       decimal.Decimal `json:"Pending"`
	CryptoAddress string          `json:"CryptoAddress"`
	Requested     bool            `json:"Requested"`
	UUID          string          `json:"Uuid"`
}

//Candle .
type Candle struct {
	TimeStamp  CandleTime      `json:"T"`
	Open       decimal.Decimal `json:"O"`
	Close      decimal.Decimal `json:"C"`
	High       decimal.Decimal `json:"H"`
	Low        decimal.Decimal `json:"L"`
	Volume     decimal.Decimal `json:"V"`
	BaseVolume decimal.Decimal `json:"BV"`
}

//NewCandles .
type NewCandles struct {
	Ticks []Candle `json:"ticks"`
}

// Currency .
type Currency struct {
	Currency        string          `json:"Currency"`
	CurrencyLong    string          `json:"CurrencyLong"`
	MinConfirmation int             `json:"MinConfirmation"`
	TxFee           decimal.Decimal `json:"TxFee"`
	IsActive        bool            `json:"IsActive"`
	CoinType        string          `json:"CoinType"`
	BaseAddress     string          `json:"BaseAddress"`
	Notice          string          `json:"Notice"`
}

// Deposit .
type Deposit struct {
	ID            int64           `json:"Id"`
	Amount        decimal.Decimal `json:"Amount"`
	Currency      string          `json:"Currency"`
	Confirmations int             `json:"Confirmations"`
	LastUpdated   jTime           `json:"LastUpdated"`
	TxID          string          `json:"TxId"`
	CryptoAddress string          `json:"CryptoAddress"`
}

// Distribution .
type Distribution struct {
	Distribution   []BalanceDist   `json:"Distribution"`
	Balances       decimal.Decimal `json:"Balances"`
	AverageBalance decimal.Decimal `json:"AverageBalance"`
}

// BalanceDist .
type BalanceDist struct {
	BalanceDist decimal.Decimal `json:"Balance"`
}

// Market .
type Market struct {
	MarketCurrency     string          `json:"MarketCurrency"`
	BaseCurrency       string          `json:"BaseCurrency"`
	MarketCurrencyLong string          `json:"MarketCurrencyLong"`
	BaseCurrencyLong   string          `json:"BaseCurrencyLong"`
	MinTradeSize       decimal.Decimal `json:"MinTradeSize"`
	MarketName         string          `json:"MarketName"`
	IsActive           bool            `json:"IsActive"`
	Notice             string          `json:"Notice"`
	IsSponsored        bool            `json:"IsSponsored"`
	LogoURL            string          `json:"LogoUrl"`
}

// MarketSummary .
type MarketSummary struct {
	MarketName     string          `json:"MarketName"`
	High           decimal.Decimal `json:"High"`
	Low            decimal.Decimal `json:"Low"`
	Ask            decimal.Decimal `json:"Ask"`
	Bid            decimal.Decimal `json:"Bid"`
	OpenBuyOrders  int             `json:"OpenBuyOrders"`
	OpenSellOrders int             `json:"OpenSellOrders"`
	Volume         decimal.Decimal `json:"Volume"`
	Last           decimal.Decimal `json:"Last"`
	BaseVolume     decimal.Decimal `json:"BaseVolume"`
	PrevDay        decimal.Decimal `json:"PrevDay"`
	TimeStamp      string          `json:"TimeStamp"`
}

// OrderSimple .
type OrderSimple struct {
	OrderUUID         string          `json:"OrderUuid"`
	Exchange          string          `json:"Exchange"`
	TimeStamp         jTime           `json:"TimeStamp"`
	OrderType         string          `json:"OrderType"`
	Limit             decimal.Decimal `json:"Limit"`
	Quantity          decimal.Decimal `json:"Quantity"`
	QuantityRemaining decimal.Decimal `json:"QuantityRemaining"`
	Commission        decimal.Decimal `json:"Commission"`
	Price             decimal.Decimal `json:"Price"`
	PricePerUnit      decimal.Decimal `json:"PricePerUnit"`
}

// Order For getorder
type Order struct {
	AccountID                  string
	OrderUUID                  string `json:"OrderUuid"`
	Exchange                   string `json:"Exchange"`
	Type                       string
	Quantity                   decimal.Decimal `json:"Quantity"`
	QuantityRemaining          decimal.Decimal `json:"QuantityRemaining"`
	Limit                      decimal.Decimal `json:"Limit"`
	Reserved                   decimal.Decimal
	ReserveRemaining           decimal.Decimal
	CommissionReserved         decimal.Decimal
	CommissionReserveRemaining decimal.Decimal
	CommissionPaid             decimal.Decimal
	Price                      decimal.Decimal `json:"Price"`
	PricePerUnit               decimal.Decimal `json:"PricePerUnit"`
	Opened                     string
	Closed                     string
	IsOpen                     bool
	Sentinel                   string
	CancelInitiated            bool
	ImmediateOrCancel          bool
	IsConditional              bool
	Condition                  string
	ConditionTarget            decimal.Decimal
}

// OrderBook .
type OrderBook struct {
	Buy  []OrderBookOrder `json:"buy"`
	Sell []OrderBookOrder `json:"sell"`
}

// OrderBookOrder .
type OrderBookOrder struct {
	Quantity decimal.Decimal `json:"Quantity"`
	Rate     decimal.Decimal `json:"Rate"`
}

// Ticker .
type Ticker struct {
	Bid  decimal.Decimal `json:"Bid"`
	Ask  decimal.Decimal `json:"Ask"`
	Last decimal.Decimal `json:"Last"`
}

//Trade is used in getmarkethistory
type Trade struct {
	OrderUUID int64           `json:"Id"`
	Timestamp jTime           `json:"TimeStamp"`
	Quantity  decimal.Decimal `json:"Quantity"`
	Price     decimal.Decimal `json:"Price"`
	Total     decimal.Decimal `json:"Total"`
	FillType  string          `json:"FillType"`
	OrderType string          `json:"OrderType"`
}

// UUID .
type UUID struct {
	ID string `json:"uuid"`
}

// Withdrawal .
type Withdrawal struct {
	PaymentUUID    string          `json:"PaymentUuid"`
	Currency       string          `json:"Currency"`
	Amount         decimal.Decimal `json:"Amount"`
	Address        string          `json:"Address"`
	Opened         jTime           `json:"Opened"`
	Authorized     bool            `json:"Authorized"`
	PendingPayment bool            `json:"PendingPayment"`
	TxCost         decimal.Decimal `json:"TxCost"`
	TxID           string          `json:"TxId"`
	Canceled       bool            `json:"Canceled"`
}
