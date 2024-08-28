package iterations

// workout/iterations       1.235s  with Builder 41332561                25.88 ns/op
// workout/iterations       1.394s  with +=character 11218129               102.0 ns/op
func Repeat(seriesCounter int, character string) string {
	//var builder strings.Builder
	var result string
	if seriesCounter == 0 {
		seriesCounter = 5
	}
	for i := 0; i < seriesCounter; i++ {
		result += character
		//builder.WriteString(character)
	}
	return result
}
