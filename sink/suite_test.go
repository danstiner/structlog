package sink

import (
	"github.com/danstiner/structlog"
	"github.com/stretchr/testify/suite"
)

type SinkTestSuite struct {
	suite.Suite
	sink structlog.Sink
}

func NewSinkTestSuite(sink structlog.Sink) *SinkTestSuite {
	return &SinkTestSuite{
		sink: sink,
	}
}

func (suite *SinkTestSuite) TestTrace() {
	suite.sink.Log(structlog.Event{
		Level: structlog.TraceLevel,
	})
}

func (suite *SinkTestSuite) TestInfo() {
	suite.sink.Log(structlog.Event{
		Level: structlog.InfoLevel,
	})
}

func (suite *SinkTestSuite) TestError() {
	suite.sink.Log(structlog.Event{
		Level: structlog.ErrorLevel,
	})
}

func (suite *SinkTestSuite) TestPanic() {
	suite.sink.Log(structlog.Event{
		Level: structlog.PanicLevel,
	})
}
