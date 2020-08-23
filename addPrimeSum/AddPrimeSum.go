package main
import ("fmt"; "os")

func myAtoi(str string) int {
	var nbr int = 0
	var i int = 0
	var neg bool = false
	if str[i] == '-' || str[i] == '+' {
		if str[i] == '-' {
			neg = true
		}
		i++
	}
	for i < len(str) && (str[i] >= '0' && str[i] <= '9') {
		nbr = nbr * 10 + (int(str[i]) - 48)
		i++
	}
	if (neg) {
		nbr = -nbr
	}
	return nbr
}

func isPrime(num int) bool {
	var count int = 0
	for i := 1; i <= num; i++ {
		if num % i == 0 {
			count++
		}
		if count > 2 {
			return false
		}
	}
	return true
}

func findPrime(end int) int {
	if end <= 1 {
		return 0
	}
	digits := make([]int, end)
	for i := 1; (i - 1) < len(digits); i++ {
		digits[i-1] = i;
	}
	digits = append(digits[:0], digits[1:] ...)
	for i := 0; i < len(digits); i++ {
		if isPrime(digits[i]) {
			for j := i + 1; j < len(digits); j++ {
				if digits[j] % digits[i] == 0 {
					digits = append(digits[:j], digits[j + 1:] ...)
				}
			}
			if i > 10 && (digits[i] * digits[i]) > end {
				break
			}
		}
	}
	return sumPrime(digits)
}

func sumPrime(prime []int) int {
	var tmp int = 0
	for i := 0; i < len(prime); i++ {
		tmp = tmp + prime[i]
	}
	return tmp
}

func main() {
	if len(os.Args) == 2 {
		fmt.Println(findPrime(myAtoi(os.Args[1])))
	}else {
		fmt.Println(0)
	}
}
