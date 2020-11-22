all: clean pacgo

pacgo:
	@echo "Building pacgo binary for use on local system..."
	@go build -o bin/pacgo ./cmd/pacgo

clean:
	@echo "Cleaning bin/..."
	@rm -rf bin/*
