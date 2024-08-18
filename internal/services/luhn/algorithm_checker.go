package luhn

import "strconv"

type luhnAlgorithmChecker struct {
}

func Checker() *luhnAlgorithmChecker {
	return &luhnAlgorithmChecker{}
}

func (c luhnAlgorithmChecker) Check(val string) bool {
	var sum int64
	nDigits := len(val)
	parity := nDigits % 2
	for i := 0; i < nDigits; i++ {
		digit, _ := strconv.ParseInt(string(val[i]), 10, 64)
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
