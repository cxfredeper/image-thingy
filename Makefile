encode decode: %: %.go $(wildcard codec/*.go)
	go build $<

.PHONY: wasm
wasm: docs/codec.wasm

docs/codec.wasm: wasm.go $(wildcard codec/*.go)
	tinygo build -no-debug -target wasm -o $@ $<
