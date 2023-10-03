package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("request-id")
	fmt.Println(rID)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Go context tut")
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	fmt.Println(ctx.Err())

	ctx = enrichContext(ctx)

	go doSomethingCool(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Exceeded timeout")
		fmt.Println(ctx.Err())
	}
	time.Sleep(2 * time.Second)
}
