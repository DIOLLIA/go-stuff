package roma

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roma   string
}{
	{Arabic: 1, Roma: "I"},
	{Arabic: 2, Roma: "II"},
	{Arabic: 3, Roma: "III"},
	{Arabic: 4, Roma: "IV"},
	{Arabic: 9, Roma: "IX"},
	{Arabic: 18, Roma: "XVIII"},
	{Arabic: 20, Roma: "XX"},
	{Arabic: 39, Roma: "XXXIX"},
	{Arabic: 40, Roma: "XL"},
	{Arabic: 47, Roma: "XLVII"},
	{Arabic: 49, Roma: "XLIX"},
	{Arabic: 50, Roma: "L"},
	{Arabic: 100, Roma: "C"},
	{Arabic: 90, Roma: "XC"},
	{Arabic: 400, Roma: "CD"},
	{Arabic: 500, Roma: "D"},
	{Arabic: 900, Roma: "CM"},
	{Arabic: 1000, Roma: "M"},
	{Arabic: 1984, Roma: "MCMLXXXIV"},
	{Arabic: 3999, Roma: "MMMCMXCIX"},
	{Arabic: 2014, Roma: "MMXIV"},
	{Arabic: 1006, Roma: "MVI"},
	{Arabic: 798, Roma: "DCCXCVIII"},
}

func TestRomanConverter(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("convert arabic %d to roma %q", test.Arabic, test.Roma), func(t *testing.T) {
			result := ConvertToRoman(test.Arabic)
			if test.Roma != result {
				t.Errorf("expected %q, got %q", test.Roma, result)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("convert roma  %q to arabic %d", test.Roma, test.Arabic), func(t *testing.T) {
			result := ConvertToArabic(test.Roma)
			if test.Arabic != result {
				t.Errorf("expected %q, got %q", test.Arabic, result)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			log.Println(arabic)
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 10}); err != nil {
		t.Error("failed checks", err)
	}
}
