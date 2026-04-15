package logger

import "go.uber.org/zap"

// 先简化，先让日志系统存在，后续再考虑日志级别、日志格式等
func New() (*zap.Logger, error) {
	return zap.NewProduction()
}
