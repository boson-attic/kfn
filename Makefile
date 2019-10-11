all: build image
.PHONY: all

PACKR := $(shell which packr)

check:
ifeq ($(PACKR),)
$(error No packr in PATH=$(PATH), consider installing https://github.com/gobuffalo/packr)
else
$(info packr: $(PACKR))
endif

gen:
	cd dsl && antlr4 -Dlanguage=Go -o ../pkg/dsl/gen -package gen -listener -visitor Kfn.g4

build: check
	packr build -v -o kfn main.go

install: check
	packr install

image: build
	buildah bud --format docker --tag kfn .
