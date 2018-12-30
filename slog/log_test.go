package slog

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoggerTestSuite struct {
	suite.Suite
	logger Logger
}

func NewLoggerTestSuite(logger Logger) *LoggerTestSuite {
	return &LoggerTestSuite{
		logger: logger,
	}
}

func (suite *LoggerTestSuite) TestTrace() {
	suite.logger.Trace("message")
}

func (suite *LoggerTestSuite) TestInfo() {
	suite.logger.Info("message")
}

func (suite *LoggerTestSuite) TestError() {
	suite.logger.Error("message")
}

func (suite *LoggerTestSuite) TestPanic() {
	suite.Panics(func() { suite.logger.Panic("message") })
}

func TestNoop(t *testing.T) {
	suite.Run(t, NewLoggerTestSuite(Noop()))
}
