.PHONY: all clean generate test

all: generate

clean:
	rm -f public/client.gen.go beta/betaclient.gen.go

generate:
	cd public && go generate ./...
	cd beta && go generate ./...
	go build -o fix-generated-code ./fix-generated-code.go
	@echo "Fixing generated code issues (thank you oapi-codegen ❤️)"
	@while go build ./public 2>&1 | grep -q "cannot use.*untyped string constant"; do \
		./fix-generated-code ./public/client.gen.go; \
	done
	@while go build ./beta 2>&1 | grep -q "cannot use.*untyped string constant"; do \
		./fix-generated-code ./beta/betaclient.gen.go; \
	done
	rm -f fix-generated-code
	go mod tidy

update-specs:
	./.github/scripts/fetch-specs.sh

update-changelogs:
	./.github/scripts/fetch-changelogs.sh