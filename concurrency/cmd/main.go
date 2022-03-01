package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/rahardiandj/go-advance/concurrency"
	"time"
)

func main() {
	ctx := context.Background()
	err := concurrency.Go(ctx, "test log", func(ctx context.Context) error {

		fmt.Printf("Test goroutine\n")
		panic("panic from goroutine")

		return errors.New("test")

	})

	if err != nil {
		fmt.Printf("Got panic=%v\n", err)
	}

	time.Sleep(1 * time.Second)
}
