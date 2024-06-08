gen:
	bash bin/allgen.sh

build: gen
	go build -o inetchkdr main.go
