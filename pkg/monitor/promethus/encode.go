package promethus

import (
	"fmt"
	"strings"
)

func encodeQuery(keys []string, querys map[string]string) string {
	quote := func(k, v string) string {
		return fmt.Sprintf("%s='%s'")
	}
	var buf strings.Builder
	for key, value := range querys {
		for _, k := range keys {
			if key == k && len(value) > 0 {
				buf.WriteString(quote(key, value))
				if buf.Len() > 0 {
					buf.WriteByte(',')
				}
			}
		}
	}
	return buf.String()
}
