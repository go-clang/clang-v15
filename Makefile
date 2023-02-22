.PHONY: all test docker/test

export CC := clang
export CXX := clang++

LLVM_LIBDIR?=$(shell llvm-config --libdir)
LLVM_VERSION?=15

GO_TEST_FUNC?=.

all: test

test:
	CGO_LDFLAGS="-L${LLVM_LIBDIR} -Wl,-rpath,${LLVM_LIBDIR}" go test -v -race -shuffle=on -run=${GO_TEST_FUNC} ./...

docker/test:
	docker container run --rm -it -e GO_TEST_FUNC=${GO_TEST_FUNC} --mount type=bind,src=$(CURDIR),dst=/go/src/github.com/go-clang/clang-v${LLVM_VERSION} -w /go/src/github.com/go-clang/clang-v${LLVM_VERSION} ghcr.io/go-clang/base:${LLVM_VERSION} make test
