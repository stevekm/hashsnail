SHELL:=/bin/bash
.ONESHELL:

format:
	gofmt -l -w .

test:
	set -euo pipefail
	go clean -testcache && \
	go test -v ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

CHARS:=abcdefghijklmnopqrstuvwxyz
test-run:
	set +e; set -x;
	go run . 0cc175b9c0f1b6a831c399e269772661 # a
	go run . 900150983cd24fb0d6963f7d28e17f72 # abc
	go run . 26ca5bfe74f8de88ccaac5c0f44b349d # abcc
	go run . ab56b4d92b40713acc5af89985d4b786 --char-set $(CHARS) # abcde
	go run . e2fc714c4727ee9395f324cd2e7f331f # abcd
	go run . ab56b4d92b40713acc5af89985d4b786 # abcde

# This one is deadlocking at {~ a415703380621ae08574dd5a1f2cb579
# go run . e2fc714c4727ee9395f324cd2e7f331f --max-size 2


# // comparison; <4s on M1 MacBook Air
# // $ ./hashcat -m 0 -a 3 ab56b4d92b40713acc5af89985d4b786
# // wantedHash := "0cc175b9c0f1b6a831c399e269772661" // a 0.225s
# // wantedHash := "900150983cd24fb0d6963f7d28e17f72" // abc 2s
# // wantedHash := "e2fc714c4727ee9395f324cd2e7f331f" // abcd 2:36
# // wantedHash := "ab56b4d92b40713acc5af89985d4b786" // abcde 1:40:50

# >>> FOUND value 'abcc' for hash 26ca5bfe74f8de88ccaac5c0f44b349d
# go run . 26ca5bfe74f8de88ccaac5c0f44b349d --threads 8  34.42s user 6.60s system 251% cpu 16.281 total

# >>> FOUND value 'a' for hash 0cc175b9c0f1b6a831c399e269772661
# go run . 0cc175b9c0f1b6a831c399e269772661 --threads 8  0.20s user 0.20s system 147% cpu 0.268 total

# >>> FOUND value 'abc' for hash 900150983cd24fb0d6963f7d28e17f72
# go run . 900150983cd24fb0d6963f7d28e17f72 --threads 8  0.59s user 0.28s system 193% cpu 0.452 total

# >>> FOUND value 'abcd' for hash e2fc714c4727ee9395f324cd2e7f331f
# go run . e2fc714c4727ee9395f324cd2e7f331f --threads 8  35.02s user 7.01s system 248% cpu 16.893 total

# >>> FOUND value 'abcde' for hash ab56b4d92b40713acc5af89985d4b786
# go run . ab56b4d92b40713acc5af89985d4b786 --threads 8 --char-set abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRST  98.44s user 20.32s system 250% cpu 47.402 total



HUNTER2:=2ab96390c7dbe3439de74d0c9b0b1767
HUNTER:=6b1b36cbb04b41490bfc0ab2bfa26f86
HUNTE:=9e3ae1b513b828922d4f691254bda0c1
# (4068424926 hashes, 1h28m58.767790988s)
HUNT:=bc9bf7bb6c4ab8d0daf374963110f4a7
# (56243782 hashes, 1m3.520669683s)
HUN:=fe1b3b54fde5b24bb40f22cdd621f5d0
# (584970 hashes, 575.490021ms)
HU:=18bd9197cb1d833bc352f47535c00320
# (7153 hashes, 5.649212ms)
H:=2510c39011c5be704182423e3a695e91
# (69 hashes, 308.426µs)

HASHES:=$(H) $(HU) $(HUN) $(HUNT) $(HUNTE) $(HUNTER) $(HUNTER2)
test-hashes:
	for i in $(HASHES); do time ./hashsnail $$i; done


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
