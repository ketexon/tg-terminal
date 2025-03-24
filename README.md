# tg-terminal

This is the terminal client for TG

## Builiding

There are 3 steps to building:

1. Codegen the proto files using `proto.sh`
2. Compiling the vfs folder into archives using `make-archive.sh`
3. Building/running the binary using `go run main.go` or `go build -o build`
