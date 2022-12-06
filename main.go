package main

import (
	"os"
	"os/exec"

	"example.com/poc-go-test/internal"
	"github.com/alecthomas/kong"
)

var cli struct {
	RunTest struct{} `cmd:"" help:"Invoke IT tool."`
}

func main() {
	internal.Log("Main CLI invoked")

	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "run-test":
		cmd := exec.Command("./bin/it", "run")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

	default:
		panic(ctx.Command())
	}
}
