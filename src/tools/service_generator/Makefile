build:
	go build -o bin/main main.go

create-service: build
	./bin/main generateBlackspaceServiceTemplate -a "REST" -s $(servicename)
	mv $(servicename) ~/go/src/github.com/Lens-Platform/Platform/src/services

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

generate: build create-service

all: build compile create-service
