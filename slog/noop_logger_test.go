package slog

import (
	"testing"
)

func TestNoopLogMethods(t *testing.T) {
	log := NewNoopLogger()
	log.Debug("message")
	log.Error("message")
	log.Info("message")
	log.Warn("message")
}

func TestNoopWith(t *testing.T) {
	log := NewNoopLogger()
	log.With("k", "v").With("k", "v")
}
