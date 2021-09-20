package rotation

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewHook(t *testing.T) {
	hook, err := NewHook(WithMaxSize(10), WithFilename("./logs/app.log"))
	require.Nil(t, err)

	logrus.AddHook(hook)

	for i := 0; i < 1000000; i++ {
		logrus.Info("logrus-rotation-hook-message-", i)
	}
}
