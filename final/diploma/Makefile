SRC_DIR = ./cmd
DEST_DIR = ./bin
NAME_SERVER = statusPage

.PHONY: all
all: clean vendor build run

.PHONY: clean
clean:
	rm -f $(DEST_DIR)/statusPage-*

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: build
build: clean vendor
	bin/go-executables-build.bash $(NAME_SERVER) $(SRC_DIR) $(DEST_DIR)

.PHONY: run
run: build
	go run -mod=mod ../skillbox-diploma/main.go 1>/dev/null &
	bin/go-executables-run.bash
