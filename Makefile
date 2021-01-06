.PHONY: all
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

run:
	go run api/main.go

test:
	go test ./...

sec:
	go get github.com/securego/gosec/v2/cmd/gosec
	gosec -exclude=G104 -tests ./...

build-mocks:
	@go get github.com/golang/mock/mockgen@v1.4.4
	@~/go/bin/mockgen -source=entity/user/interface.go -destination=mocks/user_mock.go -package=mock
	@~/go/bin/mockgen -source=pkg/password/password.go -destination=mocks/password_mock.go -package=mock
	@~/go/bin/mockgen -source=pkg/jwt/jwt.go -destination=mocks/jwt_mock.go -package=mock
