package integrationtests

import (
	"testing"

	"example.com/poc-go-test/integrationtests/suites"
	"example.com/poc-go-test/internal"
	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/suite"
)

var cli struct {
	Run struct {
		Options string
	} `cmd:"" help:"Launch integration tests."`
}

func TestMain(t *testing.T) {
	internal.Log("IT invoked")

	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "run":
		suite.Run(t, new(suites.ExampleTestSuite))

	default:
		panic(ctx.Command())
	}
}
