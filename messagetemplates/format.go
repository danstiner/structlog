package messagetemplates

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Template represents a parsed template string that can be rendered
type Template struct {
	holeCount int // number of holes for values present in the template
	length    int // length of the original template string
	parts     []part
}

type kind int

const (
	stringPart    = iota
	serializeHole = iota
	stringifyHole = iota
)

type part struct {
	argIndex int
	kind     kind

	// for a hole part, this string is the name of the hole
	// for a string part, this is the string itself
	string string
}

// Format parses a template string and then renders it with the given values
func Format(template string, values ...interface{}) (string, []KV, error) {
	parsed, err := Parse(template)
	if err != nil {
		return "", nil, fmt.Errorf("TODO: %v", err)
	}
	return Render(parsed, values...)
}

// Parse takes a template string and produces a renderable Template
func Parse(template string) (Template, error) {
	var argIndex int
	var lastIndex int
	var index int
	var holeCount int

	length := len(template)
	partsHeuristic := strings.Count(template, "{")*2 + 1
	parts := make([]part, 0, partsHeuristic)
	reader := strings.NewReader(template)

	// Template ::= ( Text | EscapedOpenBrace | Hole )*
	for reader.Len() > 0 {
		r, s, _ := reader.ReadRune()
		index += s

		if r != '{' {
			// Text ::= [^\{]
			continue
		}

		if reader.Len() == 0 {
			return Template{}, errors.New("unclosed opening brace")
		}

		parts = append(parts, part{
			kind:   stringPart,
			string: template[lastIndex : index-1],
		})
		lastIndex = index

		r, s, _ = reader.ReadRune()
		index += s
		if r == '{' {
			// EscapedOpenBrace :: = '{{'
			continue
		}

		// Hole ::=  '{' Operator? Id (',' Alignment)? (':' Format)? '}'
		hole := part{
			argIndex: argIndex,
			kind:     stringifyHole,
		}
		argIndex++

		// Operator ::= ('@' | '$')?
		if r == '@' {
			hole.kind = serializeHole
			lastIndex = index
			r, s, _ = reader.ReadRune()
			index += s
			if reader.Len() == 0 {
				return Template{}, errors.New("unexpected end of Hole")
			}
		} else if r == '$' {
			hole.kind = stringifyHole
			lastIndex = index
			r, s, _ = reader.ReadRune()
			index += s
			if reader.Len() == 0 {
				return Template{}, errors.New("unexpected end of Hole")
			}
		}

		// Id ::= (Name | Index)
		// Name ::= [0-9A-z_]+
		// Index::= [0-9]+
		var isIndex = true
		var j = index
		for reader.Len() > 0 {
			if !isNameRune(r) {
				break
			}
			isIndex = isIndex && isIndexRune(r)
			j = index
			r, s, _ = reader.ReadRune()
			index += s
		}

		if lastIndex == index {
			return Template{}, errors.New("Name/Index must have at least one character")
		}

		if r != ',' && r != ':' && r != '}' {
			var holeType string
			if isIndex {
				holeType = "Index"
			} else {
				holeType = "Name"
			}
			return Template{}, fmt.Errorf("unexpected end of %s", holeType)
		}

		hole.string = template[lastIndex:j]
		lastIndex = index

		if isIndex {
			var err error
			hole.argIndex, err = strconv.Atoi(hole.string)
			if err != nil {
				return Template{}, fmt.Errorf("TODO: %v", err)
			}
		}

		// (',' Alignment)?
		if r == ',' {
			// Alignment ::= '-'?[0-9]+

			// TODO
			return Template{}, errors.New("Alignment specifier not yet implemented")
		}

		// (':' Format)?
		if r == ':' {
			// Format ::= [^\{]+

			// TODO Format will require back-tracking
			return Template{}, errors.New("Format specifier not yet implemented")
		}

		// '}'
		if r != '}' {
			return Template{}, errors.New("unexpected character at end of Hole")
		}

		parts = append(parts, hole)
		lastIndex = index
		holeCount++
	}

	parts = append(parts, part{
		kind:   stringPart,
		string: template[lastIndex:],
	})

	return Template{holeCount, length, parts}, nil
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

func Render(template Template, values ...interface{}) (string, []KV, error) {
	if template.holeCount != len(values) {
		return "", nil, errors.Errorf("Template has %d holes for values but %d values were passed", template.holeCount, len(values))
	}

	kv := make([]KV, 0, template.holeCount)
	bldr := makeBuilder(template.length)

	for _, part := range template.parts {
		if part.kind == stringPart {
			bldr.writeString(part.string)
			continue
		}

		v := values[part.argIndex]
		kv = append(kv, KV{Key: part.string, Value: v})

		switch part.kind {
		case stringifyHole:
			switch v := v.(type) {
			case int:
				bldr.writeInt(int64(v))
			case string:
				bldr.writeString(v)
			default:
				bldr.writeString(fmt.Sprintf("%v", v))
			}
		case serializeHole:
			bytes, err := json.Marshal(v)
			if err != nil {
				return "", nil, err
			}
			bldr.write(bytes)
		default:
			return "", nil, errors.New("Unknown hole type")
		}
	}

	return bldr.string(), kv, nil
}
