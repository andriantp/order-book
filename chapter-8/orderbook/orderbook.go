package orderbook

type Side string

const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)

type OrderType string

const (
	Market OrderType = "MARKET"
	Limit  OrderType = "LIMIT"
)

type Order struct {
	ID    int
	Type  OrderType
	Side  Side
	Price float64
	Qty   int
}

type PriceLevel struct {
	Price  float64
	Orders []*Order
}

type OrderBook struct {
	Bids       map[float64]*PriceLevel
	Asks       map[float64]*PriceLevel
	orders     map[int]*Order
	eventStore EventStore
}

func NewOrderBook(store EventStore) *OrderBook {
	return &OrderBook{
		Bids:       make(map[float64]*PriceLevel),
		Asks:       make(map[float64]*PriceLevel),
		orders:     make(map[int]*Order),
		eventStore: store,
	}
}
