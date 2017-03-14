package filters

import "strings"

func Replace(f string, t string, doc string) string {
	return strings.Replace(doc, f, t, -1)
}
