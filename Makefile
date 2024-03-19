.PHONY: test
test:
	go test -v -race -buildvcs ./...

.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

.PHONY: build-http
build-http:
	go build -o=/tmp/bin/http ./cmd/http

.PHONY: run-http
run-http: build-http
	/tmp/bin/http

.PHONY: run-http/live
run-http/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build-http" --build.bin "/tmp/bin/http" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go" \
		--misc.clean_on_exit "true"
