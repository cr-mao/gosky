package job

import "context"

type Once struct {
	Work Work
}

func (job *Once) Run(ctx context.Context) error {
	handle := RecoverInterceptor(job.Work)
	return handle(ctx)
}
