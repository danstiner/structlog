package sink

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/danstiner/structlog"
	"github.com/danstiner/structlog/messagetemplates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestJsonLevel(t *testing.T) {
	const expected = structlog.InfoLevel
	const key = "$level"
	var buf bytes.Buffer
	log := NewJson(&buf)

	log.Log(structlog.Event{
		Level: expected,
	})

	var m map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	actual, ok := m[key]
	require.True(t, ok, "logged JSON object must contain a %s meta field", key)
	assert.Equal(t, expected.String(), actual)
}

func TestJsonMessage(t *testing.T) {
	const expected = "hello world"
	const key = "$message"
	var buf bytes.Buffer
	log := NewJson(&buf)

	log.Log(structlog.Event{
		Message: expected,
	})

	var m map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	actual, ok := m[key]
	require.True(t, ok, "logged JSON object must contain a %s meta field", key)
	assert.Equal(t, expected, actual)
}

func TestJsonStructData(t *testing.T) {
	type Position struct {
		Latitude, Longitude float64
	}

	var pos = Position{25, 132}
	const key = "position"
	var buf bytes.Buffer
	log := NewJson(&buf)
	expected, err := toMap(pos)
	require.NoError(t, err)

	log.Log(structlog.Event{
		Message: "message",
		Data: []messagetemplates.KV{
			messagetemplates.KV{Key: key, Value: pos},
		},
	})

	var m map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	actual, ok := m[key]
	require.True(t, ok, "logged JSON object must contain a %s meta field", key)
	assert.Equal(t, expected, actual)
}

func TestJsonTimestamp(t *testing.T) {
	now := time.Now()
	expected := now.Truncate(time.Second)
	const key = "$timestamp"
	var buf bytes.Buffer
	sink := NewJson(&buf)

	sink.Log(structlog.Event{
		Timestamp: now,
	})

	var m map[string]interface{}
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	value, ok := m[key]
	require.True(t, ok, "logged JSON object must contain a %s meta field", key)
	actual, err := time.Parse(time.RFC3339, value.(string))
	require.NoError(t, err)
	assert.WithinDuration(t, expected, actual, time.Duration(0), "Timestamp string: '%s'", value)
}

func TestJsonSink(t *testing.T) {
	var buf bytes.Buffer
	suite.Run(t, NewSinkTestSuite(NewJson(&buf)))
}

func toMap(i interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
