package sink

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

func TestLogrusSink(t *testing.T) {
	suite.Run(t, NewSinkTestSuite(LogrusSink{logrus.StandardLogger()}))
}
