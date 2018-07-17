package Iteration

const repeatCount5 = 5
const repeatCount10 = 10

func Repeat5(char string) string {
	var repeated string
	for i := 0; i < repeatCount5; i++ {
		repeated += char
	}

	return repeated
}

func Repeat10(char string) string {
	var repeated string
	for i := 0; i < repeatCount10; i++ {
		repeated += char
	}

	return repeated
}
