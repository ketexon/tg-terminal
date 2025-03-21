$env:PATH += ";" + $(go env GOPATH) + "\bin"

protoc proto/* --go_out=paths=source_relative:.