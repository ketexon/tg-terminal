$env:PATH += ";" + $(go env GOPATH) + "\bin"

protoc proto/*.proto --go_out=paths=source_relative:.
