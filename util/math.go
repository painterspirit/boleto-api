package util

import (
	"strconv"
)

type ModFunc int

const (
	MOD10 ModFunc = 1 + iota
	MOD11
)

//mod11 calculate Mod11 DV from string
func mod11(valueSequence string, baseNum ...int) int {
	base := 9
	digit := 0
	sum := 0
	weight := 2

	var values []int

	if baseNum != nil {
		base = baseNum[0]
	}

	for _, r := range valueSequence {
		c := string(r)
		n, _ := strconv.Atoi(c)
		values = append(values, n)
	}
	for i := len(values) - 1; i >= 0; i-- {
		sum += values[i] * weight

		if weight < base {
			weight = weight + 1
		} else {
			weight = 2
		}
	}
	digit = 11 - (sum % 11)
	return digit
}

//mod10 calculate Mod10 DV from string
func mod10(valueSequence string) int {
	sum := 0

	multiplyByTwo := true

	for i := len(valueSequence) - 1; i >= 0; i-- {
		c := string(valueSequence[i])

		num, _ := strconv.Atoi(c)

		if multiplyByTwo {
			num = num * 2
			sum += (num / 10) + (num % 10)
			multiplyByTwo = false
		} else {
			sum += num
			multiplyByTwo = true
		}
	}

	remainder := sum % 10

	if remainder == 0 {
		return 0
	}

	return 10 - remainder
}

//OurNumberDv calculate DV from OurNumber
func OurNumberDv(valueSequence string, modFunc ModFunc, base ...int) string {
	digit := 0

	if modFunc == MOD10 {
		digit = mod10(valueSequence)
	} else if modFunc == MOD11 && base != nil {
		digit = mod11(valueSequence, base[0])
	} else if modFunc == MOD11 {
		digit = mod11(valueSequence)
	}

	if base != nil && digit == 10 {
		return "P"
	} else if digit > 9 {
		digit = 0
	}
	return strconv.Itoa(digit)
}

//BarcodeDv calculate DV from barcode
func BarcodeDv(valueSequence string) string {
	digit := mod11(valueSequence)
	if digit <= 1 || digit > 9 {
		digit = 1
	}
	return strconv.Itoa(digit)
}
