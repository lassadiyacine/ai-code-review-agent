package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

// --- Bug 1: Division by zero (no check) ---
func divide(a, b int) int {
	return a / b
}

// --- Bug 2: Out-of-bounds access (no check) ---
func getIndex(arr []int, i int) int {
	return arr[i]
}

// --- Bug 3: SQL injection ---
func getUser(db *sql.DB, username string) (*sql.Row, error) {
	query := "SELECT * FROM users WHERE username = '" + username + "'"
	row := db.QueryRow(query)
	return row, nil
}

// --- Bug 4: Command injection ---
func runCommand(userInput string) string {
	cmd := exec.Command("sh", "-c", userInput)
	output, err := cmd.Output()
	if err != nil {
		return "error"
	}
	return string(output)
}

// --- Bug 5: Race condition ---
var counter int

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++ // no mutex, race condition
	}
}

func runRace() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementCounter(&wg)
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

// --- Bug 6: Hardcoded credentials ---
func connectDB() (*sql.DB, error) {
	password := "admin123"
	connStr := fmt.Sprintf("postgres://admin:%s@localhost:5432/marketpay?sslmode=disable", password)
	return sql.Open("postgres", connStr)
}

// --- Bug 7: Error ignored ---
func readConfig() string {
	data, _ := os.ReadFile("config.json")
	return string(data)
}

// --- Bug 8: HTTP without timeout ---
func fetchURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "error"
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	n, _ := resp.Body.Read(buf)
	return string(buf[:n])
}

// --- Bug 9: Sensitive data in logs ---
func processPayment(cardNumber string, amount float64) {
	fmt.Printf("Processing payment: card=%s amount=%.2f\n", cardNumber, amount)
}

// --- Bug 10: Resource leak (file not closed) ---
func writeLog(message string) {
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	// missing f.Close() — resource leak
	f.WriteString(message + "\n")
}
