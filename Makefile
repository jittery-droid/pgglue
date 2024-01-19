.PHONY: all test

test:
	mkdir -p bin
	go test -p 1 -count 1 \
	-coverpkg=./... -coverprofile=bin/coverage.out \
	.