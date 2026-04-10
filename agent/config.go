package agent

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	ReviewPrompt   string
	SecurityPrompt string
	SummaryPrompt  string
}

func LoadConfig() Config {
	config := Config{
		ReviewPrompt:   "Tu es un expert code. Analyse ce diff et liste uniquement les problèmes de qualité et de logique. Ignore complètement les failles de sécurité, ne les mentionne pas.\nFormat : [fichier.go] fonction: problème. Une ligne par problème, max 5.\nSi aucun problème trouvé : réponds uniquement RAS. Ne mets jamais RAS si tu as déjà listé des problèmes.",
		SecurityPrompt: "Tu es un expert sécurité. Analyse ce diff et liste uniquement les failles de sécurité.\nFormat : [fichier.go] nomFonction: description courte de la faille.\nSi aucun problème : réponds uniquement RAS.",
		SummaryPrompt:  "Résume ce diff en 2-3 phrases maximum.\nDis ce qui a changé, pas comment.",
	}

	file, err := os.Open(".review-config")
	if err != nil {
		return config
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "review_prompt":
			config.ReviewPrompt = value
		case "security_prompt":
			config.SecurityPrompt = value
		case "summary_prompt":
			config.SummaryPrompt = value
		}
	}

	return config
}
