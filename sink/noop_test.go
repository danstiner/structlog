package sink

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestNoopSink(t *testing.T) {
	suite.Run(t, NewSinkTestSuite(Noop{}))
}
