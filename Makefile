build:
	@(cd ./scripts; ./build.sh)

clean:
	rm -rf bin

run:
	@make build
	@(cd ./bin; ./taskforce)
