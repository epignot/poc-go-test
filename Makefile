build:
	go test -c ./integrationtests -o bin/it
	go build -o bin/cli