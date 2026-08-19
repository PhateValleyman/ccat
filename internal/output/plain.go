package output

type PlainRenderer struct{}

func NewPlainRenderer() *PlainRenderer {
	return &PlainRenderer{}
}

func (r *PlainRenderer) Render(text string) string {
	return text
}
