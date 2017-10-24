package main

import (
	"os"

	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/gocodic/cli/gocodic/cmd"
)

func main() {
	os.Exit(cmd.Execute(gocli.NewUI(
		gocli.Reader(os.Stdin),
		gocli.Writer(os.Stdout),
		gocli.ErrorWriter(os.Stderr),
	)).Int())
}
