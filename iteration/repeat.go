package iteration

const repeatCount = 5

func Repeat(value string) string {
	var result string

	for i := 0; i < repeatCount; i++ {
		result += value
	}
	return result
}
