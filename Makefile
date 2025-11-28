build:
	go build -o bin/tuiapp ./cmd/main.go
run:
	./bin/tuiapp
delete:
	rm ./bin/tuiapp
