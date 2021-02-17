GOCMD := go1.16

.PHONY: build
build:
	@$(GOCMD) build -ldflags="-s -w" -trimpath -o go116 main.go
