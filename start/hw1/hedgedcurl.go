package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	args := newCommand()

	if err := checkCommand(&args); err != nil {
		fmt.Printf("Arguments error: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(args.Time)*time.Second)
	defer cancel()

	resChan := make(chan string)

	hedgedcurlStart(ctx, &args, resChan)

	select {
	case firstResponse := <-resChan:
		fmt.Print(firstResponse)
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("\n Error 228: Network Timeout")
			os.Exit(228)
		}
	}
}
