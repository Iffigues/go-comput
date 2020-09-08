GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
GOINIT=$(GOCMD) mod init computor
GOU= ./computor
all:  build

build:
	$(GOFMT)
	$(GOINIT)
	$(GOBUILD)
fmt:
	$(GOFMT)
clean: 
	$(GOCLEAN)
	rm -rf $(GOU)
fclean:
	$(GOCLEAN)
	rm -rf $(GOU) go.sum go.mod
