package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	agent "github.com/lassadiyacine/ai-code-review-agent/agent"
)

func main() {
	// Review directe sans commit
	if len(os.Args) > 1 && os.Args[1] == "--file" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . --file <fichier.go>")
			os.Exit(1)
		}
		content, err := os.ReadFile(os.Args[2])
		if err != nil {
			fmt.Println("Erreur lecture fichier:", err)
			os.Exit(1)
		}
		mode := agent.GetMode()
		if mode == "cancel" {
			fmt.Println("Analyse annulée.")
			os.Exit(0)
		}
		length := agent.AskLength()
		if length == 0 {
			fmt.Println("Analyse annulée.")
			os.Exit(0)
		}
		result := agent.CallGemini(string(content), mode, length)
		fmt.Println(result)
		return
	}

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

	// Vérifier si c'est un preset
	if len(os.Args) > 1 {
		if preset, exists := agent.GetPreset(os.Args[1]); exists {
			result := agent.CallGemini(string(output), preset.Mode, preset.MaxTokens)
			fmt.Println(result)
			return
		}
	}

	mode := agent.GetMode()
	if mode == "cancel" {
		fmt.Println("Analyse annulée.")
		os.Exit(0)
	}

	length := agent.AskLength()
	if length == 0 {
		fmt.Println("Analyse annulée.")
		os.Exit(0)
	}
	result := agent.CallGemini(string(output), mode, length)
	fmt.Println(result)
}
