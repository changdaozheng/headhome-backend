.PHONY: run
run: main
	./$<

main: rm main 
	./cmd/main.go go.mod
	go build -o $@ ./cmd/main.go

.PHONY: all
all: main
