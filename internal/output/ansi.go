package output

import (
	"github.com/PhateValleyman/ccat/v2/internal/theme"
)

type ANSIRenderer struct {
	Theme *theme.Theme
}

func NewANSIRenderer(th *theme.Theme) *ANSIRenderer {
	return &ANSIRenderer{Theme: th}
}

func (r *ANSIRenderer) Render(text string) string {
	// text už obsahuje escape sekvence z highlight
	return text
}
