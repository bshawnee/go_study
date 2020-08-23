package main
import "fmt"

func isPositive(nbr int) bool {
	if nbr > 0 {
		return true
	}
	return false
}

func abc(nbr int) int {
	if isPositive(nbr) {
		return nbr
	}else if nbr == 0 {
		return 0
	}else {
		return nbr * -1
	}
}

func main() {
	fmt.Println(isPositive(1))
	fmt.Println(abc(0))
}
