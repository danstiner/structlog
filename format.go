package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type template struct {
	text  []byte
	holes []hole
}

type hole struct {
	alignment int
	argIndex  int
	name      string
	serialize bool
	stringify bool
	textIndex int
}

func parse(format string) (*template, error) {
	var buf bytes.Buffer
	var holes []hole

	reader := strings.NewReader(format)

	// Template ::= ( Text | EscapedOpenBrace | Hole )*
	for reader.Len() > 0 {
		ch, _, _ := reader.ReadRune()
		if ch != '{' {
			// Text ::= [^\{]
			buf.WriteRune(ch)
			continue
		}

		if reader.Len() == 0 {
			return nil, errors.New("unclosed opening brace")
		}

		ch, _, _ = reader.ReadRune()
		if ch == '{' {
			// EscapedOpenBrace :: = '{{'
			_, _ = buf.WriteRune('{')
			_, _ = buf.WriteRune('{')
			continue
		}

		// Hole ::=  '{' Operator? Id (',' Alignment)? (':' Format)? '}'
		i, _ := reader.Seek(0, io.SeekCurrent)
		hole := hole{
			argIndex:  len(holes),
			textIndex: buf.Len(),
		}

		// Operator ::= ('@' | '$')?
		if ch == '@' || ch == '$' {
			hole.serialize = ch == '@'
			hole.stringify = ch == '$'
			ch, _, _ = reader.ReadRune()
			if reader.Len() == 0 {
				return nil, errors.New("unexpected end of Hole")
			}
		}

		// Id ::= (Name | Index)
		// Name ::= [0-9A-z_]+
		// Index::= [0-9]+
		var id strings.Builder
		var isIndex = true
		i, _ = reader.Seek(0, io.SeekCurrent)
		for reader.Len() > 0 {
			if !isNameRune(ch) {
				break
			}
			_, _ = id.WriteRune(ch)
			isIndex = isIndex && isIndexRune(ch)
			ch, _, _ = reader.ReadRune()
		}

		j, _ := reader.Seek(0, io.SeekCurrent)

		if i == j {
			return nil, errors.New("Name/Index must have at least one character")
		}

		if ch != ',' && ch != ':' && ch != '}' {
			var holeType string
			if isIndex {
				holeType = "Index"
			} else {
				holeType = "Name"
			}
			return nil, fmt.Errorf("unexpected end of %s", holeType)
		}

		hole.name = id.String()
		i = j

		if isIndex {
			var err error
			hole.argIndex, err = strconv.Atoi(hole.name)
			if err != nil {
				return nil, fmt.Errorf("TODO: %v", err)
			}
		}

		// (',' Alignment)?
		if ch == ',' {
			// Alignment ::= '-'?[0-9]+

			// TODO
			return nil, errors.New("Alignment specifier not yet implemented")
		}

		// (':' Format)?
		if ch == ':' {
			// Format ::= [^\{]+

			// TODO Format will require back-tracking
			return nil, errors.New("Format specifier not yet implemented")
		}

		// '}'
		if ch != '}' {
			return nil, errors.New("unexpected character at end of Hole")
		}

		holes = append(holes, hole)
	}

	return &template{
		text:  buf.Bytes(),
		holes: holes,
	}, nil
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

func render(format string, v ...interface{}) (string, map[string]interface{}, error) {
	template, err := parse(format)
	if err != nil {
		return "", nil, fmt.Errorf("TODO: %v", err)
	}

	if len(template.holes) != len(v) {
		return "", nil, errors.New("TODO")
	}

	m := make(map[string]interface{})
	var bldr strings.Builder
	var textIndex int
	for i, h := range template.holes {
		m[h.name] = v[i]

		// Copy bytes from text buffer between the last hole (if any)
		// and the current hole into the output builder
		_, _ = bldr.Write(template.text[textIndex:h.textIndex])
		textIndex = h.textIndex

		// Format value of hole and write to output builder
		err := writeHoleFmt(&h, v[i], &bldr)
		if err != nil {
			return "", nil, fmt.Errorf("TODO: %v", err)
		}
	}

	// Copy remaining bytes from text buffer
	_, _ = bldr.Write(template.text[textIndex:])

	return bldr.String(), m, nil
}

func writeHoleFmt(h *hole, v interface{}, b *strings.Builder) error {
	if h.serialize {
		bytes, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("TODO: %v", err)
		}
		_, err = b.Write(bytes)
		return err
	}
	_, err := b.WriteString(fmt.Sprintf("%v", v))
	return err
}
