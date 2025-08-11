APP_NAME := imchat
BIN_DIR := bin

.PHONY: all clean linux windows mac

all: linux windows mac

linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)_linux .

windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)_windows.exe .

mac:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)_mac .

clean:
	@echo "Cleaning binaries..."
	rm -rf $(BIN_DIR)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# 自动确保在编译前创建 bin 目录
linux windows mac: | $(BIN_DIR)
