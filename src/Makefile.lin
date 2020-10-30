#help
linux:
	@echo "welcome to surin linux make file\n[*]build\n[*]run\n[*]compile-windows\n[*]compile-freebsd"

help:
	@echo "[*]build\n[*]run\n[*]compile-linux\n[*]compile-freebsd"
#run project
run:
	GOOS=linux go run main.go
	
#build project on linux
build:
	GOOS=linux go build main.go
	
#build project on windows	
compile-windows:
	@echo "Compiling for windows : SurinBuild/main-windows-386.exe"
	GOOS=windows GOARCH=386 go build -o SurinBuild/main-windows-386.exe main.go
	
#build project on freebsd
compile-freebsd:
	@echo "Compiling for freebsd : SurinBuild/main-freebsd-386"
	GOOS=freebsd GOARCH=386 go build -o SurinBuild/main-freebsd-386 main.go
	
#compile all
compile-all:
	@echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o SurinBuild/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o SurinBuild/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o SurinBuild/main-windows-386.exe main.go