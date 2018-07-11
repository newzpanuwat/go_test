package challenges

import (
	"fmt"
)

func Convert(value float64) float64 {
	return value * 37.5
}

func main() {
	fmt.Println(Convert(2))
}
