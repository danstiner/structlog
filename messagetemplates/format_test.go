package messagetemplates

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoHoles(t *testing.T) {
	msg, m, err := Format("Message")
	require.NoError(t, err)
	assert.Equal(t, "Message", msg)
	assert.Equal(t, map[string]interface{}{}, m)
}

func TestHole(t *testing.T) {
	msg, m, err := Format("{msg}", "Hello world!")
	require.NoError(t, err)
	assert.Equal(t, "Hello world!", msg)
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
		Long: 132,
	}
	msg, m, err := Format("Processed {@position} in {elapsed} ms", pos, 34)
	require.NoError(t, err)
	assert.Equal(t, `Processed {"Lat":25,"Long":132} in 34 ms`, msg)
	assert.Equal(t, map[string]interface{}{
		"position": pos,
		"elapsed":  34,
	}, m)
}
