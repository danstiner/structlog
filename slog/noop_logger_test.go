package slog

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNoopLogMethods(t *testing.T) {
	log := NewNoopLogger()
	log.Debug("message")
	log.Error("message")
	log.Info("message")
	log.Warn("message")
}

func TestNoopErrorWrap(t *testing.T) {
	log := NewNoopLogger()
	err := errors.New("error")
	werr := log.WrapError(err, "additional context")
	assert.Equal(t, err, werr, "No-op logger should not modify error")
}
