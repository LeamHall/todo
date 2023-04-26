.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint:	fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: test
	rm -rf bin
	mkdir bin
	export GOARCH=amd64 GOOS=linux;  go build -o bin/todo_$${GOOS}.$${GOARCH} main.go
	export GOARCH=arm64 GOOS=linux;	 go build -o bin/todo_$${GOOS}.$${GOARCH} main.go
	export GOARCH=amd64 GOOS=darwin; go build -o bin/todo_$${GOOS}.$${GOARCH} main.go
	export GOARCH=amd64 GOOS=windows; go build -o bin/todo_$${GOOS}.$${GOARCH} main.go
.PHONY:build

test:	vet
	go test ./...
.PHONY:test

longtest: vet
	go test -count 1000 -timeout 30m ./...
.PHONY:longtest

distro: build
	rm -rf tmp
	mkdir tmp
	cp docs/README.txt bin
	cd bin && zip -r ../tmp/todo.zip .
.PHONY:distro

clean:
	# this may be OBE
	#rm cmd/npcGen/npcGen
	#rm cmd/crewgen/crewgen
.PHONY:clean

