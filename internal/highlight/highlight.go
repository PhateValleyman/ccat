package highlight

import (
	"github.com/PhateValleyman/ccat/v2/internal/theme"
	"regexp"
)

type Highlighter struct {
	Lang   string
	Theme  *theme.Theme
	Tokens []TokenPattern
}

type TokenPattern struct {
	Type  string
	Regex *regexp.Regexp
}

func NewHighlighter(lang string, th *theme.Theme) *Highlighter {
	patterns := GetPatterns(lang)
	return &Highlighter{
		Lang:   lang,
		Theme:  th,
		Tokens: patterns,
	}
}

func (h *Highlighter) HighlightLine(line string) string {
	// Pro každý token najdeme shody a nahradíme je barevnými kódy
	// Postup: procházíme řádek a postupně aplikujeme regex, vybíráme nejdelší shody?
	// Jednodušší: aplikujeme všechny regex, ale to může způsobit překryvy.
	// Pro ukázku použijeme jednoduchou metodu: každý token obalíme stylem.
	// Použijeme metodu ReplaceAllString s funkcí.
	// Abychom se vyhnuli překryvům, použijeme techniku "procházení po značkách".
	// Zde pro stručnost implementujeme základní verzi:
	var result string
	pos := 0
	for pos < len(line) {
		bestMatch := -1
		bestToken := ""
		bestStart := -1
		bestEnd := -1
		// Najdeme nejbližší shodu
		for _, t := range h.Tokens {
			loc := t.Regex.FindStringIndex(line[pos:])
			if loc != nil {
				start := pos + loc[0]
				end := pos + loc[1]
				if bestStart == -1 || start < bestStart {
					bestStart = start
					bestEnd = end
					bestToken = t.Type
					bestMatch = 1
				}
			}
		}
		if bestMatch == -1 {
			// žádná shoda
			result += line[pos:]
			break
		}
		// přidej text před shodou
		result += line[pos:bestStart]
		// obal shodu stylem
		style := h.Theme.GetStyle(bestToken)
		colored := style.Apply(line[bestStart:bestEnd])
		result += colored
		pos = bestEnd
	}
	return result
}
