encode decode: %: %.go $(wildcard codec/*.go)
	go build $<

.PHONY: wasm
wasm: site/codec.wasm

site/codec.wasm: wasm.go $(wildcard codec/*.go)
	tinygo build -target wasm -o $@ $<
