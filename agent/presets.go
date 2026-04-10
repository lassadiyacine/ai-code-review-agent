package agent

type Preset struct {
	Mode      string
	MaxTokens int
	Label     string
}

var Presets = map[string]Preset{
	"--quick": {
		Mode:      "review",
		MaxTokens: 100,
		Label:     "Quick Review",
	},
	"--full": {
		Mode:      "review",
		MaxTokens: 600,
		Label:     "Full Review",
	},
	"--security": {
		Mode:      "security",
		MaxTokens: 300,
		Label:     "Security Review",
	},
}

func GetPreset(arg string) (Preset, bool) {
	preset, exists := Presets[arg]
	return preset, exists
}
