//+build !go1.12

package httprouter

import "strings"

func toLower(s string) string {
	return strings.ToLower(s)
}
