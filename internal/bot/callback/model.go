package callback

import "fmt"

const PatternLen = 4

type Pattern [PatternLen]byte

// NewPattern may panic
func NewPattern(s string) Pattern {
	return *(*Pattern)(([]byte(s))[:])
}

func MakeCallbackData(p Pattern, data string) string {
	return string(p[:]) + data
}

func Parse(s string) (p Pattern, data string, err error) {
	if len(s) < PatternLen {
		err = fmt.Errorf("cant get pattern from %q", s)
		return
	}

	p = *(*Pattern)([]byte(s)[:PatternLen])
	data = s[PatternLen:]

	return
}
