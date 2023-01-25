compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

build1:
	go build -o bin/main main.go

run:
	go run main.go

hello:
	echo "Thanks for installing GoTracker, wait for go to build"

build: hello build1
