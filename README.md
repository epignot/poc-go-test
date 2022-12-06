# poc-go-test

## Goal

Demonstrate: 
- We can build a test binary with `go test` using the `testing` module
- Invoke it through a regular binary
- Co-locate the code in the same repo

## Instructions 
- Install deps : `go mod tidy`
- Build binaries : `make build`
- Invoke the integration tool directly: `./bin/it run`
- Or through the "main cli" binary: `./bin/cli run-test`

## Caveats

I'm not able to run the default `go test` command without errors right now,   
`go test ./integrationtests` passes extra args and kong fails trying to parse them