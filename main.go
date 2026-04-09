package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	agent "github.com/lassadiyacine/ai-code-review-agent/agent"
)

func main() {
	cmd := exec.Command("git", "diff", "--staged")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Erreur git diff: %v", err)
		os.Exit(1)
	}

	// Si lancé depuis le hook
	if len(os.Args) > 1 && os.Args[1] == "--hook" {
		if !agent.AskReview() {
			fmt.Println("Review ignorée, commit en cours...")
			os.Exit(0)
		}
	}

	mode := agent.GetMode()
	result := agent.CallGemini(string(output), mode)
	fmt.Println(result)
}
