package highlight

import (
	"regexp"
)

var patternsCache = make(map[string][]TokenPattern)

func GetPatterns(lang string) []TokenPattern {
	if p, ok := patternsCache[lang]; ok {
		return p
	}

	var p []TokenPattern
	switch lang {
	case "go":
		p = []TokenPattern{
			{Type: "keyword", Regex: regexp.MustCompile(`\b(break|case|chan|const|continue|default|defer|else|fallthrough|for|func|go|goto|if|import|interface|map|package|range|return|select|struct|switch|type|var)\b`)},
			{Type: "string", Regex: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
			{Type: "comment", Regex: regexp.MustCompile(`//.*$`)},
			{Type: "number", Regex: regexp.MustCompile(`\b\d+(\.\d+)?\b`)},
		}
	case "python":
		p = []TokenPattern{
			{Type: "keyword", Regex: regexp.MustCompile(`\b(and|as|assert|break|class|continue|def|del|elif|else|except|finally|for|from|global|if|import|in|is|lambda|nonlocal|not|or|pass|raise|return|try|while|with|yield)\b`)},
			{Type: "string", Regex: regexp.MustCompile(`(?:r|u|f)?'(?:[^'\\]|\\.)*'|(?:r|u|f)?"(?:[^"\\]|\\.)*"`)},
			{Type: "comment", Regex: regexp.MustCompile(`#.*$`)},
			{Type: "number", Regex: regexp.MustCompile(`\b\d+(\.\d+)?\b`)},
		}
	case "markdown":
		p = []TokenPattern{
			{Type: "keyword", Regex: regexp.MustCompile(`(?m)^#{1,6}\s.*$`)},     // headers
			{Type: "string", Regex: regexp.MustCompile("`[^`]+`")},             // inline code
			{Type: "comment", Regex: regexp.MustCompile(`(?m)^>\s.*$`)},         // blockquote
			{Type: "number", Regex: regexp.MustCompile(`\[[^\]]+\]\([^\)]+\)`)}, // links
		}
	case "bash":
		p = []TokenPattern{
			{Type: "keyword", Regex: regexp.MustCompile(`\b(if|then|else|elif|fi|for|while|in|do|done|case|esac|function|local|return)\b`)},
			{Type: "string", Regex: regexp.MustCompile(`'(?:[^'\\]|\\.)*'|"(?:[^"\\]|\\.)*"`)},
			{Type: "comment", Regex: regexp.MustCompile(`#.*$`)},
			{Type: "number", Regex: regexp.MustCompile(`\b\d+\b`)},
		}
	case "json":
		p = []TokenPattern{
			{Type: "keyword", Regex: regexp.MustCompile(`\b(true|false|null)\b`)},
			{Type: "string", Regex: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
			{Type: "number", Regex: regexp.MustCompile(`-?\b\d+(\.\d+)?([eE][+-]?\d+)?\b`)},
		}
	case "yaml":
		p = []TokenPattern{
			{Type: "keyword", Regex: regexp.MustCompile(`^\s*[\w-]+\s*:`)}, // keys
			{Type: "string", Regex: regexp.MustCompile(`'(?:[^'\\]|\\.)*'|"(?:[^"\\]|\\.)*"`)},
			{Type: "comment", Regex: regexp.MustCompile(`#.*$`)},
			{Type: "number", Regex: regexp.MustCompile(`\b\d+(\.\d+)?\b`)},
		}
	default:
		p = []TokenPattern{
			{Type: "string", Regex: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
			{Type: "comment", Regex: regexp.MustCompile(`//.*$|/\*.*?\*/`)},
		}
	}
	patternsCache[lang] = p
	return p
}

