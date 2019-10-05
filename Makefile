build:
	@(go build -o "bin/taskforce" github.com/restechnica/taskforce/cmd/taskforce)

clean:
	rm -rf bin

install:
	@(go install github.com/restechnica/taskforce/cmd/taskforce)

run:
	@make build
	@(cd ./bin; ./taskforce)
