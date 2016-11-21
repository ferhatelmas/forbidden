// Package forbidden is a handy collection of strings
// to prevent users from signup with.
package forbidden

import "strings"

var nameMap map[string]bool

func init() {
	nameMap = make(map[string]bool, len(names))
	for _, name := range names {
		nameMap[name] = true
	}
}

// Name checks if given name is forbidden.
func Name(name string) bool {
	return nameMap[strings.ToLower(strings.TrimSpace(name))]
}
