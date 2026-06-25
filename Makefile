SERVICES = tree dobf preprocess func

define compile
	$(GO) build \
	-ldflags "-s -w" \
	-gcflags "all=-N -l" \
	-o bin/$(1) cmd/$(1)/main.go
endef

all: $(SERVICES)

proto:
	$(BUF) generate

$(SERVICES):
	$(call compile,$(@))

GO        		 := CGO_ENABLED=0 GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) go
BUF              := go tool github.com/bufbuild/buf/cmd/buf