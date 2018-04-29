package slog

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJsonDebug(t *testing.T) {
	var buf bytes.Buffer
	log := NewJsonLogger(&buf)

	log.Debug("message")

	var m map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	assert.Equal(t, map[string]interface{}{
		"@level":    "DEBUG",
		"@message":  "message",
		"@template": "message",
	}, m)
}

func TestJsonError(t *testing.T) {
	var buf bytes.Buffer
	log := NewJsonLogger(&buf)

	log.Error("message")

	var m map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	assert.Equal(t, map[string]interface{}{
		"@level":    "ERROR",
		"@message":  "message",
		"@template": "message",
	}, m)
}

func TestJsonInfo(t *testing.T) {
	var buf bytes.Buffer
	log := NewJsonLogger(&buf)

	log.Info("message")

	var m map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	assert.Equal(t, map[string]interface{}{
		"@level":    "INFO",
		"@message":  "message",
		"@template": "message",
	}, m)
}

func TestJsonWarn(t *testing.T) {
	var buf bytes.Buffer
	log := NewJsonLogger(&buf)

	log.Warn("message")

	var m map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	assert.Equal(t, map[string]interface{}{
		"@level":    "WARN",
		"@message":  "message",
		"@template": "message",
	}, m)
}

func TestJsonWith(t *testing.T) {
	var buf bytes.Buffer
	log := NewJsonLogger(&buf)
	log.With("k", "v").With("k", "v")
}
