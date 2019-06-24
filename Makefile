GO_CMD=go
PANDORA_BIN=pandora_bin
PANDORA_BIN_PATH=./output/bin
PANDORA_ETC_PATH=./output/etc
PANDORA_PUBLIC_PATH=./output/public
PANDORA_LOG_PATH=./output/log
PANDORA_RES_PATH=./output/resource

.PHONY: all
all: clean build install

.PHONY: linux
linux: clean build-linux install

.PHONY: build
build:
	@echo "build pandora start >>>"
	GOPROXY=https://goproxy.io $(GO_CMD) mod tidy
	$(GO_CMD) build -o $(PANDORA_BIN) ./pandora/main.go
	@echo ">>> build pandora complete"

.PHONY: install
install:
	@echo "install pandora start >>>"
	mkdir -p $(PANDORA_BIN_PATH)
	mv $(PANDORA_BIN) $(PANDORA_BIN_PATH)/pandora
	mkdir -p $(PANDORA_ETC_PATH)
	cp ./pandora.ini $(PANDORA_ETC_PATH)
	cp -r ./public $(PANDORA_PUBLIC_PATH)
	mkdir -p $(PANDORA_LOG_PATH)
	cp -r ./resource $(PANDORA_RES_PATH)
	@echo ">>> install pandora complete"

.PHONY: clean
clean:
	@echo "clean start >>>"
	rm -fr ./output
	rm -f $(PANDORA_BIN)
	@echo ">>> clean complete"

.PHONY: build-linux
build-linux:
	@echo "build-linux start >>>"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -o $(PANDORA_BIN) ./pandora/main.go
	@echo ">>> build-linux complete"

.PHONY: fmt
fmt:
	@echo "format code >>>"
	find . -name '*.go' | grep -v vendor | xargs gofmt -w -s
	@echo ">>> format code completed"
