GOCMD := go1.16rc1

.PHONY: build
build:
	@$(GOCMD) build -ldflags="-s -w" -trimpath -o go116 main.go
