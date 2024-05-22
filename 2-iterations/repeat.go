package iterations

const defaultRepeatCount = 5

func Repeat(character string, repeatCount int) string {
	if repeatCount <= 0 {
		repeatCount = defaultRepeatCount
	}
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
