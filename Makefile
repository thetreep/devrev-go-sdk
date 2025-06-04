.PHONY: all clean generate test

all: generate

clean:
	rm -f public/client.gen.go beta/betaclient.gen.go

generate:
	cd public && go generate ./...
	cd beta && go generate ./...
	go mod tidy

update-specs:
	./.github/scripts/fetch-specs.sh

update-changelogs:
	./.github/scripts/fetch-changelogs.sh