package main

import (
	"syscall"

	"okcoding.com/gardnr/internal/cli"
	"okcoding.com/gardnr/internal/cmdutil"
)

func main() {
	ctx, cancel := cmdutil.ContextWithSignal(syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	cli.Execute(ctx)
}
