all:
	make build && make run
debug:
	make build && DEBUG=1 make run
build:
	go build -o bin/mangareadertui ./cmd/tui/main.go
run:
	./bin/mangareadertui
delete:
	rm ./bin/mangareadertui
