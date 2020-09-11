package main

import "fmt"

func algorithmLuhn(s string) bool {
	num := 0
	sum := 0
	second := len(s) % 2
	for i := len(s) - 1; i >= 0; i-- {
		num = int(s[i] - '0')
		if i%2 == second {
			num *= 2
			if num > 9 {
				num = num%10 + 1
			}
		}
		sum += num
	}
	return sum%10 == 0
}

func main() {
	fmt.Println(algorithmLuhn("79927398713")) // true
	fmt.Println(algorithmLuhn("79927398710")) // false
}
