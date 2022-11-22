package iteration

func Repeat(chart string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + chart

	}
	return repeated
}
