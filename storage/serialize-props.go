package storage

import (
	"strings"
)

func serializeProps(prefix string, props []string) string {
	out := []string{}
	for _, s := range props {
		out = append(out, prefix+s)
	}

	return strings.Join(out, ",")
}
