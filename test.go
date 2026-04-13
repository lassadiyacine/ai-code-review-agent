package main

func multiply(a int, b int) int {
	return a * b
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func getIndex(arr []int, i int) int {
	return arr[i]
}

func buildSQL(table string, id string) string {
	return "SELECT * FROM " + table + " WHERE id = " + id
}
