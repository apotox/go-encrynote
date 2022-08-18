.PHONY: build clean

go_apps = bin/create_note bin/read_note bin/delete_note

bin/%:
	@echo building the lambda functions... $(*F)
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o $@ handlers/$(*F)/$(*F).go

build: $(go_apps) | vendor

tests: clean build
	@echo running tests...
	go test -v -timeout 30s ./functions/...

localtests: clean build localenv
	@echo running tests...
	go test -v -timeout 30s ./encryption/...
	go test -v -timeout 30s ./functions/...

localenv:
	docker compose -f ./tools/docker-compose.yaml up -d

clean:
	go clean -testcache
	rm -f $(go_apps) Gopkg.lock

vendor: go.mod
	go mod tidy
	go mod vendor