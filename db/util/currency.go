package util

//Consts for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	RUB = "RUB"
	SUM = "SUM"
)

//IsSupportedCurrency return true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, RUB, SUM:
		return true
	}
	return false
}
