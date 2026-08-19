package output

import (
	"github.com/PhateValleyman/ccat/v2/internal/theme"
	"html"
)

type HTMLRenderer struct {
	Theme *theme.Theme
}

func NewHTMLRenderer(th *theme.Theme) *HTMLRenderer {
	return &HTMLRenderer{Theme: th}
}

func (r *HTMLRenderer) Render(text string) string {
	// Zde bychom měli převést barevné značky na HTML span
	// Pro jednoduchost vrátíme escapovaný text s CSS třídami
	// Toto je zjednodušené – ideálně bychom tokenizovali a obalili span.
	escaped := html.EscapeString(text)
	return "<pre>" + escaped + "</pre>"
}
