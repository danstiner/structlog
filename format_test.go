package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoHoles(t *testing.T) {
	s, m, err := render("Message")
	require.NoError(t, err)
	assert.Equal(t, s, "Message")
	assert.Equal(t, m, map[string]interface{}{})
}

func TestHole(t *testing.T) {
	s, m, err := render("Hello {msg}!", "world")
	require.NoError(t, err)
	assert.Equal(t, s, "Hello world!")
	assert.Equal(t, m, map[string]interface{}{
		"msg": "world",
	})
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
	assert.Equal(t, s, `Processed {"Lat":25,"Long":134} in 34 ms`)
	assert.Equal(t, m, map[string]interface{}{
		"position": pos,
		"elapsed":  34,
	})
}
