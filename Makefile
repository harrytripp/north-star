build:
	@mkdir -p bin
	go build -o bin/northstar ./cmd/northstar

run: build
	./bin/northstar

clean:
	rm -rf bin/