package output

type Renderer interface {
	Render(text string) string
}
