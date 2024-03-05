package cache

import (
	"context"
	"time"
)

const timeout = 5

func contextTimeout() (ctx context.Context, cancel context.CancelFunc) {
	ctx = context.Background()
	ctx, cancel = context.WithTimeout(ctx, timeout*time.Second)
	return
}
