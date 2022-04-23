.PHONY: build

run:build start

start:
	./.bin/app

build:
	go build -o .bin/app -v cmd/web/main.go

# windows -------------------------------------------------------------------------
build_exe_64:
	env GOOS=windows GOARCH=amd64 go build -o .bin/app_win64.exe -v cmd/web/main.go

build_exe_32:
	env GOOS=windows GOARCH=386 go build -o .bin/app_win32.exe -v cmd/web/main.go

# TODO: доделать make версий
# linux ---------------------------------------------------------------------------
# build_linux_64:
# 	env GOOS=windows GOARCH=amd64 go build -o .bin/app_win64 -v cmd/web/main.go

# build_linux_32:
# 	env GOOS=windows GOARCH=amd64 go build -o .bin/app_win32 -v cmd/web/main.go

# macos ---------------------------------------------------------------------------
# build_mac_64:
# 	env GOOS=windows GOARCH=amd64 go build -o .bin/app_win64 -v cmd/web/main.go

# build_mac_32:
# 	env GOOS=windows GOARCH=amd64 go build -o .bin/app_win32 -v cmd/web/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := run