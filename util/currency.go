package util

// const for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency{
	case USD, EUR, CAD:
		return true
	default:
		return false
	}
}