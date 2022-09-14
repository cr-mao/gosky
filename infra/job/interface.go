package job

import "context"

type Runner interface {
	// Work(context.Context) error
	Run(context.Context) error
}

var _ Runner = (*Once)(nil)
var _ Runner = (*Daemon)(nil)
