SHELL:=/bin/bash
.ONESHELL:

format:
	gofmt -l -w .

test:
	set -euo pipefail
	go clean -testcache && \
	go test -v ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

test-run:
	set +e; set -x;
	go run . 0cc175b9c0f1b6a831c399e269772661
	go run . 900150983cd24fb0d6963f7d28e17f72
	go run . ab56b4d92b40713acc5af89985d4b786 --max-size 6
	go run . e2fc714c4727ee9395f324cd2e7f331f --max-size 2


# // comparison; <4s on M1 MacBook Air
# // $ ./hashcat -m 0 -a 3 ab56b4d92b40713acc5af89985d4b786
# // wantedHash := "0cc175b9c0f1b6a831c399e269772661" // a 0.225s
# // wantedHash := "900150983cd24fb0d6963f7d28e17f72" // abc 2s
# // wantedHash := "e2fc714c4727ee9395f324cd2e7f331f" // abcd 2:36
# // wantedHash := "ab56b4d92b40713acc5af89985d4b786" // abcde 1:40:50



build:
	go build -o ./hashsnail main.go
.PHONY:build

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