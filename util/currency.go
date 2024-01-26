package util

const (
	USD = "USD"
	GBP = "GBP"
	EUR = "EUR"
	GHS = "GHS"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, GBP, EUR, GHS:
		return true
	}
	return false
}
