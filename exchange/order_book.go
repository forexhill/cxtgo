package exchange

// OrderBook is a definition for an orderbook
type OrderBook interface {
	Symbol()
	Get()
	Head()
}
