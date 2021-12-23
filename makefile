# Go parameters
GoCMD=go
GoBUILD=$(GoCMD) build
GoCLEAN=$(GoCMD) clean
GoTEST=$(GoCMD) test
GoMOD=$(GoCMD) mod
BINARY_NAME=raw

all: build
build:
	$(GoBUILD) -o $(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
deps:
	$(GoMOD) tidy