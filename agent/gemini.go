package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func CallGemini(diff string, mode string) string {
	if mode == "cancel" {
		return "Analyse annulée."
	}
	if diff == "" {
		return "Aucun changement détecté."
	}
	if len(diff) > 10000 {
		return "Diff trop volumineux pour être analysé (> 10000 caractères). Committez en plus petites parties."
	}
	diff = FilterDiff(diff)
	if diff == "" {
		return "Aucun fichier pertinent à analyser."
	}
	if os.Getenv("GEMINI_API_KEY") == "" {
		return "Erreur : variable d'environnement GEMINI_API_KEY non définie."
	}

	var prompt string
	switch mode {
	case "security":
		prompt = `Tu es un expert sécurité. Analyse ce diff et liste uniquement les failles de sécurité.
Format : [fichier.go] nomFonction: description courte de la faille.
Si aucun problème : réponds uniquement RAS, ne mets jamais RAS si tu as déjà listé des problèmes. Ne cherche pas à tout prix des problèmes.`

	case "summary":
		prompt = `Résume ce diff en 2-3 phrases maximum.
Dis ce qui a changé, pas comment.`

	default: // review
		prompt = `Tu es un expert code. Analyse ce diff et liste uniquement les problèmes de qualité et de logique.
Ne mentionne jamais les failles de sécurité, même pour dire que tu les ignores.
Format strict : [fichier.go] fonction: problème. Une ligne par problème, max 5.
Si aucun problème trouvé : réponds uniquement le mot RAS, rien d'autre, ne mets jamais RAS si tu as déjà listé des problèmes. Ne cherche pas à tout prix des problèmes.`
	}
	prompt += "\n\n" + diff

	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{"text": prompt},
				},
			},
		},
		"generationConfig": map[string]interface{}{
			"maxOutputTokens": 100,
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "Erreur encoding JSON: " + err.Error()
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=" + apiKey

	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return "Erreur création requête: " + err.Error()
	}

	req.Header.Set("content-type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Erreur appel API: " + err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Sprintf("Erreur API (code %d): %s", resp.StatusCode, string(bodyBytes))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Erreur lecture réponse: " + err.Error()
	}

	return ParseResponse(respBody)
}
