package pbt

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	num      uint16
	expected string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	t.Run("decipher the nums", func(t *testing.T) {
		for _, test := range cases {
			tn := fmt.Sprintf("should decipher '%d' to %q", test.num, test.expected)

			t.Run(tn, func(t *testing.T) {
				got := ConvertToRoman(test.num)

				if got != test.expected {
					t.Errorf("got %q, want %q", got, test.expected)
				}
			})
		}
	})

	t.Run("decipher roman to arabic", func(t *testing.T) {
		for _, test := range cases[:4] {
			tn := fmt.Sprintf("should decipher %q to '%d'", test.expected, test.num)

			t.Run(tn, func(t *testing.T) {
				got := ConvertToArabic(test.expected)

				if got != test.num {
					t.Errorf("got %d, want %d", got, test.num)
				}
			})
		}
	})

	t.Run("run property based test", func(t *testing.T) {
		assertion := func(num uint16) bool {
			if num > 3999 {
				return true
			}
			t.Log("testing", num)
			roman := ConvertToRoman(num)
			arabic := ConvertToArabic(roman)

			return arabic == num
		}

		qc := &quick.Config{
			MaxCount: 1000,
		}

		if err := quick.Check(assertion, qc); err != nil {
			t.Error("failed checks", err)
		}
	})
}

func BenchmarkConvertToRoman(b *testing.B) {
	b.Run("decipher cases to roman", func(b *testing.B) {
		for _, nums := range cases {
			for i := 0; i < b.N; i++ {
				ConvertToRoman(nums.num)
			}
		}
	})
}

func BenchmarkConvertToArabic(b *testing.B) {
	b.Run("decipher cases to arabic", func(b *testing.B) {
		for _, nums := range cases {
			for i := 0; i < b.N; i++ {
				ConvertToArabic(nums.expected)
			}
		}
	})
}
