package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "114514")
}

func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("request-id")
	fmt.Println(rID)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time out")
			return
		default:
			fmt.Println("Doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Go Context Demo")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx = enrichContext(ctx)
	go doSomethingCool(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Oh no, I've exceeded the deadline")
		fmt.Println(ctx.Err())
	}
}
