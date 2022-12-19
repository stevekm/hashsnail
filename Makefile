SHELL:=/bin/bash
.ONESHELL:

format:
	gofmt -l -w .

test:
	set -euo pipefail
	go clean -testcache && \
	go test -v ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

# build:
# 	go build -o ./dupefinder cmd/main.go
# .PHONY:build

# # https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
# GIT_TAG:=$(shell git describe --tags)
# build-all:
# 	mkdir -p build ; \
# 	for os in darwin linux windows; do \
# 	for arch in amd64 arm64; do \
# 	output="build/dupefinder-v$(GIT_TAG)-$$os-$$arch" ; \
# 	if [ "$${os}" == "windows" ]; then output="$${output}.exe"; fi ; \
# 	echo "building: $$output" ; \
# 	GOOS=$$os GOARCH=$$arch go build -o "$${output}" cmd/main.go ; \
# 	done ; \
# 	done