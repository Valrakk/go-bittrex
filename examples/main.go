package main

import (
	"fmt"

	".."
)

func main() {
	updateChannel := make(chan bittrex.ExchangeState, 16)
	errCh := make(chan error)

	bt := bittrex.New("", "")

	go func() {
		for st := range updateChannel {
			if len(st.Fills) > 0 {
				for _, i := range st.Fills {
					fmt.Printf("Fecha: %v, Tipo: %v, Precio %v, Cantidad: %v\n", i.Timestamp, i.OrderType, i.Rate, i.Quantity)
				}
			}
		}
	}()

	go func() {
		errCh <- bt.SubscribeExchangeUpdate("BTC-ETH", updateChannel, nil)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			fmt.Println("Error")
			return
		}
	}

}
