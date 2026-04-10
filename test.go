package main

import "fmt"

func applyDiscount(price float64, discount float64) float64 {
	return price - (price * discount)
}

func safeDivide(a int, b int) int {
	return a / b
}

func authenticate(username string, password string) bool {
	if username == "admin" && password == "1234" {
		return true
	}
	return false
}

func average(numbers []int) float64 {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return float64(total / len(numbers))
}

func formatMessage(user string, msg string) string {
	return fmt.Sprintf("<b>%s</b>: %s", user, msg)
}
