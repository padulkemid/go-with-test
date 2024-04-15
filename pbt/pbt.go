package pbt

import (
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(num uint16) string {
	var result strings.Builder

	for _, roms := range allRomanNumerals {
		for num >= roms.Value {
			result.WriteString(roms.Symbol)
			num -= roms.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var total uint16 = 0

	for _, nums := range allRomanNumerals {
		for strings.HasPrefix(roman, nums.Symbol) {
			total += nums.Value
			roman = strings.TrimPrefix(roman, nums.Symbol)
		}
	}

	return total
}
