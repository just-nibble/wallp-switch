build:
	go build -o=./build/switcher main.go
install:
	mv ./build/switcher ~/.local/bin/switcher && rmdir ./build