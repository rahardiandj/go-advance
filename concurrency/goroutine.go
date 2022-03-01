package concurrency

import (
	"context"
	"fmt"
)

func Go(ctx context.Context, logTag string, f func(ctx2 context.Context) error) <-chan error {
	errCh := make(chan error, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("gconcurrent:%v", r)
				}

				errCh <- err
			}
		}()

		err := f(ctx)

		if err != nil {
			errCh <- err
		}
	}()

	return errCh
}
