package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func render(template string, v ...interface{}) (string, map[string]interface{}, error) {
	var buf []rune
	ln := len(template)
	r := []rune(template)
	argsByIndex := false
	argIndex := 0
	m := make(map[string]interface{})

	// Template ::= ( Text | Hole )*
	for i := 0; i < len(r); i++ {
		if r[i] != '{' {
			// [^\{] of:
			// Text ::= ([^\{] | '{{' | '}}')+
			buf = append(buf, r[i])
			continue
		}

		// Skip over current '{' rune
		i++
		if i >= ln {
			return "", nil, errors.New("ambiguous unclosed opening brace")
		}

		if r[i] == '{' {
			// '{{' of:
			// Text ::= ([^\{] | '{{' | '}}')+
			buf = append(buf, '{', '{')
			continue
		}

		// Hole ::=  '{' ('@' | '$')? (Name | Index) (',' Alignment)? (':' Format)? '}'
		var serialize bool
		var stringify bool
		var isIndex bool
		var name string
		var index = argIndex
		var alignment int
		// TODO Format will require back-tracking
		argIndex++

		// ('@' | '$')?
		if r[i] == '@' || r[i] == '$' {
			serialize = r[i] == '@'
			stringify = r[i] == '$'
			i++
			if i >= ln {
				return "", nil, errors.New("unexpected end of Hole")
			}
		}

		// (Name | Index)
		// Name ::= [0-9A-z_]+
		// Index::= [0-9]+
		j := i
		for ; j < len(r) && isNameRune(r[j]); j++ {
		}

		isIndex = true
		for _, c := range r[i:j] {
			isIndex = isIndex && isIndexRune(c)
		}

		if i == j {
			return "", nil, errors.New("Name/Index must has at least one character")
		}

		if r[j] != ',' && r[j] != ':' && r[j] != '}' {
			var holeType string
			if isIndex {
				holeType = "Index"
			} else {
				holeType = "Name"
			}
			return "", nil, fmt.Errorf("unexpected end of %s", holeType)
		}

		name = string(r[i:j])
		i = j

		if isIndex {
			var err error
			index, err = strconv.Atoi(name)
			if err != nil {
				return "", nil, fmt.Errorf("TODO: %v", err)
			}
		}

		// (',' Alignment)?
		if r[i] == ',' {
			// Alignment ::= '-'?[0-9]+

			// TODO
		}

		// (':' Format)?
		if r[i] == ':' {
			// Format ::= [^\{]+

			// TODO
		}

		// '}'
		if r[i] != '}' {
			return "", nil, errors.New("unexpected character at end of Hole")
		}

		if argsByIndex {
			// TODO
			continue
		}

		if index > len(v) {
			return "", nil, errors.New("Out of bounds access of args")
		}
		m[name] = v[index]
		// format and append value
		h, err := fmtHole(v[index], serialize, stringify, alignment)
		if err != nil {
			return "", nil, fmt.Errorf("TODO: %v", err)
		}
		buf = append(buf, h...)
	}

	if argIndex != len(v) {
		return "", nil, errors.New("did not match all args")
	}

	return string(buf), m, nil
}

// Name ::= [0-9A-z_]+
func isNameRune(c rune) bool {
	switch {
	case c >= 'a' && c <= 'z':
		return true
	case c >= 'A' && c <= 'Z':
		return true
	case c >= '0' && c <= '9':
		return true
	case c == '_':
		return true
	default:
		return false
	}
}

// Index::= [0-9]+
func isIndexRune(c rune) bool {
	return c >= '0' && c <= '9'
}

func fmtHole(v interface{}, serialize, stringify bool, alignment int) ([]rune, error) {
	if serialize {
		bytes, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("TODO: %v", err)
		}
		return []rune(string(bytes)), nil
	}
	return []rune(fmt.Sprintf("%v", v)), nil
}
