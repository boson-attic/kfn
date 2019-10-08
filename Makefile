all: build image
.PHONY: all

check:
	ifeq (, $(shell which packr))
	$(error "No packr in $(PATH), consider installing https://github.com/gobuffalo/packr")
	endif

build:
	packr build -v -o kfn main.go

install:
	packr install

image:
	buildah bud --format docker --tag kfn .
