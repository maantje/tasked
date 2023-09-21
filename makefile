dev:
	./bin/air
lint:
	docker run -t --rm -v $(shell pwd):/app -v ~/.cache/golangci-lint/v1.54.2:/root/.cache -w /app golangci/golangci-lint:v1.54.2 golangci-lint run -v