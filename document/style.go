package document

import (
	"strings"
)

type Style map[string]interface{}

func ParseStyle(text string) (Style, []error) {
	errs := []error{}
	style := make(Style)
	// TODO: Implement/use a proper parser
	text = strings.Trim(text, " \n\t")
	if len(text) == 0 {
		return style, nil
	}
	sentences := strings.Split(text, ";")
	for _, sentence := range sentences {
		pairs := strings.Split(sentence, ":")
		for i, t := range pairs {
			pairs[i] = strings.Trim(t, "\n\t ")
		}
		if len(pairs) < 2 {
			continue
		}
		k := strings.ToLower(pairs[0])
		v := strings.ToLower(strings.Join(pairs[1:], ":"))
		if k == "" {
			continue
		}
		if k == "display" {
			display, err := displayStyle(v)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			style[k] = display

			continue
		}
		style[k] = v
	}
	return style, errs
}

func (s Style) Equal(s2 Style) bool {
	if len(s) != len(s2) {
		return false
	}
	for k, v := range s {
		if s2[k] != v {
			return false
		}
	}

	return true
}
