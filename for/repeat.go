package iteration

import "fmt"

// Repeat returns character a specified number of times
func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}

	return repeated
}

func ExampleRepeat() {
	repeat := Repeat("b", 4)
	fmt.Println(repeat)
	// Output: "bbbb"
}
