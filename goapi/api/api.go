package api

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// Success code, usually 200.
	Code int

	// Account balance.
	Balance int64
}

type Error struct {
	// Error code.
	Code int

	// Error message.
	Message string
}
