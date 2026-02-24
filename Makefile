.PHONY: run build generate test clean

generate:
	buf generate

build: generate
	GOOS=js GOARCH=wasm go build -o static/client.wasm ./cmd/client
	rm -f static/wasm_exec.js
	cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" static/

run: build
	go run ./cmd/server

test:
	go test ./...

clean:
	rm -f static/client.wasm static/wasm_exec.js
