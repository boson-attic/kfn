all: build image
.PHONY: all

check:
	ifeq (, $(shell which packr))
	$(error "No packr in $(PATH), consider installing https://github.com/gobuffalo/packr")
	endif

build:
	cd dsl && antlr4 -Dlanguage=Go -o ../pkg/dsl/gen -package gen -listener -visitor Kfn.g4
	packr build -v -o kfn main.go

install:
	packr install

image:
	buildah bud --format docker --tag kfn .
