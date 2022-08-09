package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestLog(t *testing.T) {
	Debug("Debug", zap.String("test", "test"))
	Info("Debug", zap.String("test", "test"))
	Warn("Debug", zap.String("test", "test"))
	Error("Debug", zap.String("test", "test"))
}
