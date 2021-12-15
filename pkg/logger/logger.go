package logger

import (
	"go.uber.org/zap"
)

/* Initialize zap logger
* with New Production Mode
 */
func Initialize() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
