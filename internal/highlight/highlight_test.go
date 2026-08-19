package highlight

import (
	"github.com/PhateValleyman/ccat/v2/internal/theme"
	"strings"
	"testing"
)

func TestHighlightLine(t *testing.T) {
	th := theme.DefaultTheme()
	hl := NewHighlighter("go", th)

	line := `func main() { fmt.Println("hello") } // comment`
	highlighted := hl.HighlightLine(line)

	if !strings.Contains(highlighted, "\033[34mfunc\033[0m") {
		t.Errorf("Expected highlighted keyword 'func', got: %q", highlighted)
	}
	if !strings.Contains(highlighted, "\033[32m\"hello\"\033[0m") {
		t.Errorf("Expected highlighted string '\"hello\"', got: %q", highlighted)
	}
	if !strings.Contains(highlighted, "\033[36m// comment\033[0m") {
		t.Errorf("Expected highlighted comment '// comment', got: %q", highlighted)
	}
}

func TestHighlightMarkdown(t *testing.T) {
	th := theme.DefaultTheme()
	hl := NewHighlighter("markdown", th)

	line := "# Header"
	highlighted := hl.HighlightLine(line)

	if !strings.Contains(highlighted, "\033[34m# Header\033[0m") {
		t.Errorf("Expected highlighted markdown header, got: %q", highlighted)
	}
}

func TestHighlightBash(t *testing.T) {
	th := theme.DefaultTheme()
	hl := NewHighlighter("bash", th)

	line := `if [ $a == 1 ]; then echo "hi"; fi # comment`
	highlighted := hl.HighlightLine(line)

	if !strings.Contains(highlighted, "\033[34mif\033[0m") {
		t.Errorf("Expected highlighted bash keyword 'if', got: %q", highlighted)
	}
	if !strings.Contains(highlighted, "\033[32m\"hi\"\033[0m") {
		t.Errorf("Expected highlighted bash string '\"hi\"', got: %q", highlighted)
	}
	if !strings.Contains(highlighted, "\033[36m# comment\033[0m") {
		t.Errorf("Expected highlighted bash comment '# comment', got: %q", highlighted)
	}
}

func TestHighlightJSON(t *testing.T) {
	th := theme.DefaultTheme()
	hl := NewHighlighter("json", th)

	line := `{"key": "value", "num": 123, "bool": true}`
	highlighted := hl.HighlightLine(line)

	if !strings.Contains(highlighted, "\033[32m\"key\"\033[0m") {
		t.Errorf("Expected highlighted json string '\"key\"', got: %q", highlighted)
	}
	if !strings.Contains(highlighted, "\033[35m123\033[0m") {
		t.Errorf("Expected highlighted json number '123', got: %q", highlighted)
	}
	if !strings.Contains(highlighted, "\033[34mtrue\033[0m") {
		t.Errorf("Expected highlighted json keyword 'true', got: %q", highlighted)
	}
}

