package main
import "fmt"

func ActiveBits(n int) int {
	var count int = 0
	if n < 0 {
		n = -n
	}
	for i := 64; i != -1; i--{
		if n >> i & 1 + 0 == 1 {
			count++
		}
	}
	return count
}

func main() {
	var num int = -3
	fmt.Printf("%b\t %d", num, ActiveBits(num))
}
