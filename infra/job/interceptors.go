package job

import (
	"context"
	"fmt"

	"gosky/infra/logger"
)

type Work func(context.Context) error

func ErrorHandler(v interface{}) error {
	if err, ok := v.(error); ok {
		logger.WarnString("recover", "err_msg", err.Error())
		return err
	}
	returnErr := fmt.Errorf("Unknown Error, type: %T, value: %v", v, v)
	logger.WarnString("recover", "err_msg", returnErr.Error())
	return returnErr
}

func RecoverInterceptor(next Work) Work {
	return func(ctx context.Context) (returnErr error) {
		defer func() {
			if r := recover(); r != nil {
				returnErr = ErrorHandler(r)
			}
		}()
		return next(ctx)
	}
}
