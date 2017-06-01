.PHONY: all thriftrw

THRIFTRW_VERSION := v1.3
THRIFT_SRC := idl/dosa/dosa.thrift
OUT := .gen
export PATH := $(GOPATH)/bin:$(PATH)

THRIFTRW := thriftrw
YARPC_PLUGIN := thriftrw-plugin-yarpc
YARPC := yarpc

all:
ifeq (, $(shell command -v $(THRIFTRW) 2> /dev/null))
				go get 'go.uber.org/thriftrw'
endif
ifeq (, $(shell command -v $(YARPC_PLUGIN) 2> /dev/null))
				go get 'go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc'
endif
	@if $(THRIFTRW) --version | fgrep -vq $(THRIFTRW_VERSION); then echo "Your thriftrw is not version $(THRIFTRW_VERSION). Refusing to generate new code. If you intended to change the thriftrw version, please change the makefile."; exit 1; fi
	$(THRIFTRW) --plugin=$(YARPC) $(THRIFT_SRC) --out=$(OUT)
