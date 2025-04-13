package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {
	ctx1, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ctx, cancel := context.WithTimeout(ctx1, time.Second*2)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sleep", "10")
	err := cmd.Run()
	if err != nil {
		if ctx.Err() != nil {
			fmt.Printf("ctx error:%v\n", ctx.Err())
		} else {
			fmt.Printf("run error:%v\n", err)
		}
		os.Exit(1)
	}
}
