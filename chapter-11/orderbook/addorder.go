package orderbook

func (ob *OrderBook) AddOrder(order *Order) {
	ob.insertOrder(order)

	copied := *order

	ob.eventStore.Append(OrderPlaced{
		Order: &copied,
	})
}
