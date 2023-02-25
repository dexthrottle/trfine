

local_up:
	docker-compose -f local.yml up

local_build:
	docker-compose -f local.yml build


prod_up:
	docker-compose -f production.yml up

prod_build:
	docker-compose -f production.yml build


build_app_linux:
	env GOOS=linux GOARCH=386 go build -o bin/linux/rpine_linux_386 -v cmd/web/main.go
	env GOOS=linux GOARCH=arm go build -o bin/linux/rpine_linux_arm -v cmd/web/main.go
	env GOOS=linux GOARCH=arm64 go build -o bin/linux/rpine_linux_arm64 -v cmd/web/main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/rpine_linux_amd64 -v cmd/web/main.go

build_app_windows:
	env GOOS=windows GOARCH=amd64 go build -o bin/windows/rpine_win_64.exe -v cmd/web/main.go
	env GOOS=windows GOARCH=386 go build -o bin/windows/rpine_win_32.exe -v cmd/web/main.go

# build_android:
# 	env GOOS=android GOARCH=arm go build -o .bin/android/rpine_android_arm -v cmd/web/main.go

# build_macos:
# 	env GOOS=darwin GOARCH=386 go build -o .bin/macos/rpine_macos_386 -v cmd/web/main.go
# 	env GOOS=darwin GOARCH=amd64 go build -o .bin/macos/rpine_macos_amd64 -v cmd/web/main.go
# 	env GOOS=darwin GOARCH=arm go build -o .bin/macos/rpine_macos_arm -v cmd/web/main.go
# 	env GOOS=darwin GOARCH=arm64 go build -o .bin/macos/rpine_macos_arm64 -v cmd/web/main.go

build_all_platform:build_linux build_windows  #build_android  build_macos
	

.DEFAULT_GOAL := local_up