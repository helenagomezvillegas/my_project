package iteration

const repeatCount = 5

func Repeat(chart string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += chart

	}
	return repeated
}
