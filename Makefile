build:
	go build -o bin/challenge

run: build
	./challenge < input.txt

clean:
	rm challenge*