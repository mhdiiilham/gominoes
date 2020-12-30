run:
	go run api/main.go

test:
	go test ./...

sec:
	go get github.com/securego/gosec/v2/cmd/gosec
	gosec -exclude=G104 -tests ./...