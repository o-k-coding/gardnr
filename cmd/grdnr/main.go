package main

import (
	"syscall"

	"okcoding.com/grdnr/internal/cli"
	"okcoding.com/grdnr/internal/cmdutil"
)

func main() {
	ctx, cancel := cmdutil.ContextWithSignal(syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	cli.Execute(ctx)
}
