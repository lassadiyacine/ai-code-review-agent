package agent

import "os"

func GetMode() string {
	if len(os.Args) < 2 {
		return AskMode()
	}

	switch os.Args[1] {
	case "--hook":
		return AskMode()
	case "--security":
		return "security"
	case "--review":
		return "review"
	case "--summary":
		return "summary"
	case "--file", "--dir":
		return AskMode()
	default:
		return "review"
	}
}
