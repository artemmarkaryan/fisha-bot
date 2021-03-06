package format

import "fmt"

func Bold(text string) string { return "<b>" + text + "</b>" }

func Link(text, url string) string {
	return fmt.Sprintf(`<a href="%s">%s</a>`, url, text)
}
func Phone(text, url string) string {
	return fmt.Sprintf(`<a href="tel:%s">%s</a>`, url, text)
}
