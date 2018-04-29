package slog

import (
	"bytes"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestJsonLogMethods(t *testing.T) {
	var buf bytes.Buffer
	log := NewJsonLogger(&buf)
	log.Debug("message")
	log.Error("message")
	log.Info("message")
	log.Warn("message")
}

func TestJsonErrorWrap(t *testing.T) {
	var buf bytes.Buffer
	log := NewJsonLogger(&buf)
	err := errors.New("original error")
	werr := log.WrapError(err, "additional context")
	assert.Equal(t, err, errors.Cause(werr), "original cause should be preserved")
	assert.Regexp(t, "additional context", werr.Error(), "wrap message should be added")
}
