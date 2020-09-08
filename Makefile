GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
GOINIT=$(GOCMD) mod init computor
GOU= ./computor
all:  build
init:
	$(GOFMT)
	$(GOINIT)

build:
	$(GOFMT)
	$(GOGET)
	$(GOINSTALL)
	$(GOBUILD)
fmt:
	$(GOFMT)
clean: 
	$(GOCLEAN)
	rm -rf $(GOU)
fclean:
	$(GOCLEAN)
	rm -rf $(GOU) go.sum go.mod
