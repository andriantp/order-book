package orderbook

type Event interface {
	EventType() string
}
type OrderPlaced struct {
	Order *Order
}

func (e OrderPlaced) EventType() string {
	return "OrderPlaced"
}

type TradeExecuted struct {
	BuyOrderID  int
	SellOrderID int
	Price       float64
	Quantity    int
}

func (e TradeExecuted) EventType() string {
	return "TradeExecuted"
}

type OrderCancelled struct {
	OrderID int
	Price   float64
	Side    Side
}

func (e OrderCancelled) EventType() string {
	return "OrderCancelled"
}
