.PHONY: install
install: rtag
	mv rtag ~/.local/bin/rtag

rtag: rtag.go
	go build
