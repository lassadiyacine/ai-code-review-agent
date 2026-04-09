package agent

import "encoding/json"

type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content Content `json:"content"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

func ParseResponse(body []byte) string {
	var geminiResp GeminiResponse
	err := json.Unmarshal(body, &geminiResp)
	if err != nil {
		return "Erreur parsing JSON: " + err.Error()
	}
	if len(geminiResp.Candidates) == 0 {
		return "Aucune réponse de l'agent."
	}
	return geminiResp.Candidates[0].Content.Parts[0].Text
}
