package slog

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestJsonTrace(t *testing.T) {
	var buf bytes.Buffer
	log := Json(&buf)

	log.Trace("message")

	var m map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	assert.Equal(t, map[string]interface{}{
		"@level":    "TRACE",
		"@message":  "message",
		"@template": "message",
	}, m)
}

// func TestJsonError(t *testing.T) {
// 	var buf bytes.Buffer
// 	log := NewJsonLogger(&buf)

// 	log.Error("message")

// 	var m map[string]interface{}
// 	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
// 	assert.Equal(t, map[string]interface{}{
// 		"@level":    "ERROR",
// 		"@message":  "message",
// 		"@template": "message",
// 	}, m)
// }

// func TestJsonInfo(t *testing.T) {
// 	var buf bytes.Buffer
// 	log := NewJsonLogger(&buf)

// 	log.Info("message")

// 	var m map[string]interface{}
// 	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
// 	assert.Equal(t, map[string]interface{}{
// 		"@level":    "INFO",
// 		"@message":  "message",
// 		"@template": "message",
// 	}, m)
// }

// func TestJsonWarn(t *testing.T) {
// 	var buf bytes.Buffer
// 	log := NewJsonLogger(&buf)

// 	log.Warn("message")

// 	var m map[string]interface{}
// 	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
// 	assert.Equal(t, map[string]interface{}{
// 		"@level":    "WARN",
// 		"@message":  "message",
// 		"@template": "message",
// 	}, m)
// }

// func TestJsonWith(t *testing.T) {
// 	var buf bytes.Buffer
// 	log := NewJsonLogger(&buf)

// 	log.With("k", "v").Warn("message")

// 	var m map[string]interface{}
// 	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
// 	assert.Equal(t, map[string]interface{}{
// 		"@level":    "WARN",
// 		"@message":  "message",
// 		"@template": "message",
// 		"k":         "v",
// 	}, m)
// }

// func TestJsonSerializeHole(t *testing.T) {
// 	pos := struct {
// 		Lat  float32 `json:"lat"`
// 		Long float32 `json:"lng"`
// 	}{
// 		Lat:  25,
// 		Long: 132,
// 	}
// 	var buf bytes.Buffer
// 	log := NewJsonLogger(&buf)

// 	log.Info("Processed {@position} in {elapsed} ms", pos, 34)

// 	var m map[string]interface{}
// 	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
// 	assert.Equal(t, map[string]interface{}{
// 		"@level":    "INFO",
// 		"@message":  `Processed {"lat":25,"lng":132} in 34 ms`,
// 		"@template": "Processed {@position} in {elapsed} ms",
// 		"elapsed":   34.0,
// 		"position": map[string]interface{}{
// 			"lat": 25.0,
// 			"lng": 132.0,
// 		},
// 	}, m)
// }

func TestJson(t *testing.T) {
	var buf bytes.Buffer
	suite.Run(t, NewLoggerTestSuite(Json(&buf)))
}
