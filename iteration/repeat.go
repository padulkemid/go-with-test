package iteration

func Repeat(char string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += char
	}

	return repeated
}
