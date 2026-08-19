package theme

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Theme struct {
	Name     string
	Keyword  string
	String   string
	Number   string
	Comment  string
	Function string
	Type     string
	// další tokeny...
}

type Style struct {
	Foreground string
	Background string
	Bold       bool
}

func (t *Theme) GetStyle(tokenType string) Style {
	var fg string
	switch tokenType {
	case "keyword":
		fg = t.Keyword
	case "string":
		fg = t.String
	case "number":
		fg = t.Number
	case "comment":
		fg = t.Comment
	case "function":
		fg = t.Function
	case "type":
		fg = t.Type
	default:
		fg = ""
	}
	return Style{Foreground: fg}
}

func (s Style) Apply(text string) string {
	if s.Foreground == "" {
		return text
	}
	// ANSI escape kódy
	// Zde bychom převedli barvu na ANSI kód
	// Pro jednoduchost použijeme mapu názvů barev na kódy
	return "\033[" + colorCode(s.Foreground) + "m" + text + "\033[0m"
}

func colorCode(name string) string {
	colors := map[string]string{
		"black":   "30",
		"red":     "31",
		"green":   "32",
		"yellow":  "33",
		"blue":    "34",
		"magenta": "35",
		"cyan":    "36",
		"white":   "37",
		"default": "39",
	}
	if code, ok := colors[name]; ok {
		return code
	}
	return "39"
}

func Load(name string) (*Theme, error) {
	if name == "" {
		name = "default"
	}
	themesDir := filepath.Join(os.Getenv("HOME"), ".config", "ccat", "themes")
	path := filepath.Join(themesDir, name+".toml")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// fallback: embedded themes
		return DefaultTheme(), nil
	}
	var t Theme
	if _, err := toml.DecodeFile(path, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func DefaultTheme() *Theme {
	return &Theme{
		Name:     "default",
		Keyword:  "blue",
		String:   "green",
		Number:   "magenta",
		Comment:  "cyan",
		Function: "yellow",
		Type:     "red",
	}
}
