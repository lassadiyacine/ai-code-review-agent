package agent

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func AskReview() bool {
	fmt.Print("Voulez-vous une review avant de commit ? (o/n) > ")
	input := readLine()
	return input == "o" || input == "O"
}

func AskMode() string {
	fmt.Println("Quel type d'analyse ?")
	fmt.Println("1. Review de code")
	fmt.Println("2. Sécurité")
	fmt.Println("3. Résumé")
	fmt.Println("4. Annuler")
	fmt.Print("> ")

	switch readLine() {
	case "1":
		return "review"
	case "2":
		return "security"
	case "3":
		return "summary"
	case "4":
		return "cancel"
	default:
		return "review"
	}
}

func AskLength() int {
	fmt.Println("Quelle longueur de réponse ?")
	fmt.Println("1. Courte")
	fmt.Println("2. Moyenne")
	fmt.Println("3. Longue")
	fmt.Println("4. Annuler")
	fmt.Print("> ")

	switch readLine() {
	case "1":
		return 100
	case "2":
		return 300
	case "3":
		return 600
	case "4":
		return 0
	default:
		return 100
	}
}
