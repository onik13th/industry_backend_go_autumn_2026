package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}

	ctxErr := ctx.Err()
	if ctxErr != nil {
		return nil, ctxErr
	}
	if len(in) == 0 {
		return []Result[R]{}, nil
	}

	results := make([]Result[R], len(in))
	jobs := make(chan int)
	var wg sync.WaitGroup

	for range min(workers, len(in)) {
		wg.Go(func() {
			for index := range jobs {
				if ctx.Err() != nil {
					return
				}

				value, err := fn(ctx, in[index])
				if err != nil {
					results[index].Err = err
					continue
				}

				results[index].Value = value
			}
		})
	}

	dispatchJobs(ctx, len(in), jobs)
	wg.Wait()

	ctxErr = ctx.Err()
	if ctxErr != nil {
		return nil, ctxErr
	}

	return results, nil
}

func dispatchJobs(ctx context.Context, count int, jobs chan<- int) {
	defer close(jobs)

	for index := range count {
		if ctx.Err() != nil {
			return
		}

		select {
		case <-ctx.Done():
			return
		case jobs <- index:
		}
	}
}
