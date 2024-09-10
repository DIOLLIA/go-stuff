package roma

import "strings"

type RomaNumber struct {
	Arabic uint16
	Roma   string
}

var allRomaNumbers = []RomaNumber{
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

func ConvertToRoman(numb uint16) string {
	var resultNumber = strings.Builder{}

	for _, numeral := range allRomaNumbers {
		for numb >= numeral.Arabic {
			resultNumber.WriteString(numeral.Roma)
			numb -= numeral.Arabic
		}
	}

	return resultNumber.String()
}

func ConvertToArabic(romaNum string) uint16 {
	var arabic uint16 = 0
	for _, val := range allRomaNumbers {
		for strings.HasPrefix(romaNum, val.Roma) {
			arabic += val.Arabic
			romaNum = romaNum[len(val.Roma):]
		}
	}
	return arabic
}
