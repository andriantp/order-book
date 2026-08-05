package orderbook

type Side string

const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)

type Order struct {
	ID    int
	Side  Side
	Price float64
	Qty   int
}

type PriceLevel struct {
	Price  float64
	Orders []*Order
}

type OrderBook struct {
	Bids map[float64]*PriceLevel
	Asks map[float64]*PriceLevel
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		Bids: make(map[float64]*PriceLevel),
		Asks: make(map[float64]*PriceLevel),
	}
}
