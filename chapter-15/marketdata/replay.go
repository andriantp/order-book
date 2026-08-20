package marketdata

import "order-book/orderbook"

func (rm *ReadModel) Replay(events []orderbook.Event) {
	for _, event := range events {
		switch e := event.(type) {

		case orderbook.OrderPlaced:
			rm.PlaceOrder(
				e.Order.Side,
				e.Order.Price,
				e.Order.Qty,
			)

		case orderbook.TradeExecuted:
			rm.ExecuteTrade(
				e.Price,
				e.Quantity,
			)

		case orderbook.OrderCancelled:
			switch e.Side {
			case orderbook.Buy:
				delete(rm.Bids, e.Price)
				rm.RefreshBestBid()

			case orderbook.Sell:
				delete(rm.Asks, e.Price)
				rm.RefreshBestAsk()
			}

		}
	}
}
