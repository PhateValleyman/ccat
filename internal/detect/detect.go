package detect

import (
	"bufio"
	"io"
	"path/filepath"
	"strings"
)

func Detect(name string, r io.Reader) string {
	ext := filepath.Ext(name)
	switch ext {
	case ".go":
		return "go"
	case ".py":
		return "python"
	case ".sh", ".bash":
		return "bash"
	case ".json":
		return "json"
	case ".yaml", ".yml":
		return "yaml"
	case ".rs":
		return "rust"
	case ".c":
		return "c"
	case ".cpp", ".cc", ".cxx":
		return "cpp"
	case ".md":
		return "markdown"
	case ".toml":
		return "toml"
	case ".xml":
		return "xml"
	case ".js":
		return "javascript"
	case ".ts":
		return "typescript"
	case ".lua":
		return "lua"
	case ".php":
		return "php"
	case ".sql":
		return "sql"
	}
	// Shebang detection
	br := bufio.NewReader(r)
	firstLine, err := br.ReadString('\n')
	if err == nil && strings.HasPrefix(firstLine, "#!") {
		shebang := strings.Fields(firstLine)
		if len(shebang) > 1 {
			interp := filepath.Base(shebang[1])
			switch interp {
			case "python", "python3", "python2":
				return "python"
			case "bash", "sh":
				return "bash"
			case "node":
				return "javascript"
			case "lua":
				return "lua"
			case "php":
				return "php"
			}
		}
	}
	// Obsahová detekce by mohla být složitější, pro ukázku vracíme "unknown"
	return "unknown"
}
