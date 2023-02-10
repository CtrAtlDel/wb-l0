.PHONY: build
client:
	go build -v ivankvasov/project/cmd/main
	go run ivankvasov/project/cmd/main
publisher:
	go build ivankvasov/publisher/cmd
	go run ivankvasov/publisher/cmd
.DEFAULT_GOAL := client