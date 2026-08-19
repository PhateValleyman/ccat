package app

import (
	"github.com/PhateValleyman/ccat/v2/internal/config"
	"github.com/PhateValleyman/ccat/v2/internal/detect"
	"github.com/PhateValleyman/ccat/v2/internal/highlight"
	"github.com/PhateValleyman/ccat/v2/internal/output"
	"github.com/PhateValleyman/ccat/v2/internal/pager"
	"github.com/PhateValleyman/ccat/v2/internal/security"
	"github.com/PhateValleyman/ccat/v2/internal/theme"
	"github.com/PhateValleyman/ccat/v2/internal/utils"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	showLineNumbers bool
	colorMode       string
	forceLanguage   string
	themeName       string
	htmlOutput      bool
	showVersion     bool
)

func init() {
	flag.BoolVar(&showLineNumbers, "n", false, "Zobrazit čísla řádků")
	flag.StringVar(&colorMode, "color", "auto", "Použití barev: auto, never, always")
	flag.StringVar(&forceLanguage, "language", "", "Vynutit jazyk pro zvýraznění")
	flag.StringVar(&themeName, "theme", "", "Použít motiv (jméno souboru bez přípony)")
	flag.BoolVar(&htmlOutput, "html", false, "Výstup ve formátu HTML")
	flag.BoolVar(&showVersion, "version", false, "Zobrazit verzi")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ccat v2.0 – chytrý cat s barevným zvýrazněním\n")
		fmt.Fprintf(os.Stderr, "Použití: ccat [přepínače] [soubor...]\n")
		flag.PrintDefaults()
	}
}

func Run(args []string) error {
	flag.Parse()

	if showVersion {
		fmt.Println("ccat v2.0")
		return nil
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Překrytí konfigurace přepínači
	if themeName != "" {
		cfg.Theme = themeName
	}
	if colorMode != "" {
		switch colorMode {
		case "never":
			cfg.Color = false
		case "always":
			cfg.Color = true
		case "auto":
			// zůstane podle cfg
		}
	}
	if showLineNumbers {
		cfg.LineNumbers = true
	}
	if htmlOutput {
		cfg.HTML = true
	}

	// Načtení motivu
	th, err := theme.Load(cfg.Theme)
	if err != nil {
		return err
	}

	// Seznam souborů
	files := flag.Args()
	if len(files) == 0 {
		// čtení ze stdin
		return processInput(os.Stdin, "<stdin>", cfg, th, forceLanguage)
	}

	// Zpracování každého souboru
	for _, fname := range files {
		info, err := os.Stat(fname)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return fmt.Errorf("%s: je adresář", fname)
		}
		if !security.IsSafeFile(fname, cfg.MaxSize) {
			return fmt.Errorf("%s: binární nebo příliš velký soubor", fname)
		}

		file, err := os.Open(fname)
		if err != nil {
			return err
		}
		defer file.Close()

		if err := processInput(file, fname, cfg, th, forceLanguage); err != nil {
			return err
		}
	}
	return nil
}

func processInput(r io.Reader, name string, cfg *config.Config, th *theme.Theme, forcedLang string) error {
	// Přečteme celý obsah do paměti (s limitem)
	content, err := io.ReadAll(io.LimitReader(r, cfg.MaxSize))
	if err != nil {
		return err
	}

	// Detekce jazyka
	var lang string
	if forcedLang != "" {
		lang = forcedLang
	} else {
		lang = detect.Detect(name, strings.NewReader(string(content)))
	}

	// Připravíme si writer (výstup)
	var writer io.Writer
	var closer func() error

	if cfg.Pager && !cfg.HTML && utils.IsTerminal() {
		// Použijeme pager
		p, err := pager.StartPager()
		if err != nil {
			return err
		}
		writer = p.Stdin
		closer = p.Close
	} else {
		writer = os.Stdout
		closer = func() error { return nil }
	}
	defer closer()

	// Vytvoříme renderer
	var renderer output.Renderer
	if cfg.HTML {
		renderer = output.NewHTMLRenderer(th)
	} else if cfg.Color {
		renderer = output.NewANSIRenderer(th)
	} else {
		renderer = output.NewPlainRenderer()
	}

	// Pro ukázku použijeme jednoduchý regex-based highlighter.
	hl := highlight.NewHighlighter(lang, th)

	// Čtení řádek
	scanner := utils.NewLineScanner(strings.NewReader(string(content)))
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		// Ošetření dlouhých řádků
		if len(line) > cfg.MaxLineLength {
			line = line[:cfg.MaxLineLength] + "…"
		}
		// Zvýraznění
		highlighted := hl.HighlightLine(line)
		rendered := renderer.Render(highlighted)
		// Výstup
		if cfg.LineNumbers {
			rendered = fmt.Sprintf("%6d  %s", lineNum, rendered)
		}
		if _, err := fmt.Fprintln(writer, rendered); err != nil {
			return err
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

