package action

import (
	"strings"
)

type Action string

const (
	delimiter        = ':'
	wildcard         = '*'
	Any       Action = "*"
)

func (a Action) MatchingActions() []Action {
	if !a.IsValid() {
		return nil
	}

	var res []Action // Preallocate space depending on whether the Action is a wildcard (non-wildcard requires one extra item for the Action itself)
	if a[len(a)-1] == wildcard {
		res = make([]Action, 0, strings.Count(string(a), string(delimiter))+1)
	} else {
		res = append(make([]Action, 0, strings.Count(string(a), string(delimiter))+2), a)
	}

	// Cut Action from the right up to the last : delimiter until no delimiters remain
	for i := strings.LastIndex(string(a), string(delimiter)); i >= 0; i = strings.LastIndex(string(a), string(delimiter)) {
		res = append(res, Action(string(a)[:i+1]+string(wildcard)))
		a = a[:i]
	}

	return append(res, Any)
}

func (a Action) MatchingActionsString() []string {
	if !a.IsValid() {
		return nil
	}

	var res []string // Preallocate space depending on whether the Action is a wildcard (non-wildcard requires one extra item for the Action itself)
	if a[len(a)-1] == wildcard {
		res = make([]string, 0, strings.Count(string(a), string(delimiter))+1)
	} else {
		res = append(make([]string, 0, strings.Count(string(a), string(delimiter))+2), string(a))
	}

	// Cut Action from the right up to the last : delimiter until no delimiters remain
	for i := strings.LastIndex(string(a), string(delimiter)); i >= 0; i = strings.LastIndex(string(a), string(delimiter)) {
		res = append(res, string(a)[:i+1]+string(wildcard))
		a = a[:i]
	}

	return append(res, "*")
}

// IsValid returns whether an Action is well-formed.
func (a Action) IsValid() bool {
	if len(a) == 0 {
		return false
	}

	const minDelimiters, maxDelimiters = 1, 3 // One to three ':' delimiters allowed in an Action except the blanket wildcard

	delimiters, lastDelimiter := 0, true

	for i := range a {
		switch a[i] {
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '-':
			lastDelimiter = false
		case delimiter:
			delimiters++
			if lastDelimiter || delimiters > maxDelimiters {
				return false
			}

			lastDelimiter = true
		case wildcard:
			if i < len(a)-1 || !lastDelimiter {
				// Wildcard is only allowed in the end and following a delimiter, except in the blanket action wildcard "*"
				return false
			}

			return true
		default:
			return false
		}
	}

	return !lastDelimiter && delimiters > minDelimiters
}
