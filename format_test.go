package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoHoles(t *testing.T) {
	s, m, err := render("Message")
	require.NoError(t, err)
	assert.Equal(t, "Message", s)
	assert.Equal(t, map[string]interface{}{}, m)
}

func TestHole(t *testing.T) {
	s, m, err := render("{msg}", "Hello world!")
	require.NoError(t, err)
	assert.Equal(t, "Hello world!", s)
	assert.Equal(t, map[string]interface{}{
		"msg": "Hello world!",
	}, m)
}

func TestSerializeHole(t *testing.T) {
	pos := struct {
		Lat  float32
		Long float32
	}{
		Lat:  25,
		Long: 134,
	}
	s, m, err := render("Processed {@position} in {elapsed} ms", pos, 34)
	require.NoError(t, err)
	assert.Equal(t, `Processed {"Lat":25,"Long":134} in 34 ms`, s)
	assert.Equal(t, map[string]interface{}{
		"position": pos,
		"elapsed":  34,
	}, m)
}
