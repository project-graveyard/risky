// package model defines all the models needed for risky
package model

// TradeInfo represents the data used for trade simulation
type TradeInfo struct {
	// Capital is the amount of money available in a trading account before
	// taking a open positions
	Capital float64

	// Leverage is the leverage ratio provided by a brokerage. E.g. 2000:1
	Leverage int

	// RiskPercent is the percentage of the capital one wants to risk
	RiskPercent float64

	// TradeCount is the number of trades to be accounted for in the simulation
	TradeCount int

	// RR is the Risk/Reward ratio of each trade. E.g. 3:1
	RR float64

	// Asset is the type of security being traded
	Asset Pair
}

// Pair represents the type of trading pair
type Pair int

const (
	// XAUUSD is XAU/USD trading pair
	XAUUSD Pair = iota // 0

	// EURUSD is EUR/USD trading pair
	EURUSD // 1

	// BTCUSD is BTC/USD trading pair
	BTCUSD // 2

	// ETHUSD is ETH/USD trading pair
	ETHUSD // 3
)
