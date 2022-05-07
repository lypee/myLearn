package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	test1(ctx)
}

func test1(ctx context.Context) {
	d := time.Now().Add(time.Second * 5)
	sonCtx, cancelFunc := context.WithDeadline(ctx, d)
	ticker := time.NewTicker(time.Second)
	go func() {
		time.Sleep(2 * time.Second)
		cancelFunc()
	}()
	for {
		select {
		case <-sonCtx.Done():
			log.Println("ok")
			return
		case <- ticker.C:
			log.Println("1s...")
		default:
		}
	}
}

func manualCancel(cancelFunc context.CancelFunc) {
	cancelFunc()
}
