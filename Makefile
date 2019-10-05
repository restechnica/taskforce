build:
	@(cd ./scripts; ./build.sh)

clean:
	rm -rf bin

run:
	@(cd ./bin; ./taskforce)
