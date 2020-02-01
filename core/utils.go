package core

func divmod(numerator, denominator uint) (quotient, remainder uint) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}