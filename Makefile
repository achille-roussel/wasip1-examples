
files.go = $(wildcard *.go)
files.wasm = $(files.go:.go=.wasm)

all: $(files.wasm)

clean:
	rm -f $(files.wasm)

%.wasm: %.go
	GOARCH=wasm GOOS=wasip1 go build -o $@ $<
