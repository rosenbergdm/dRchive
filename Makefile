######################################################################
# @author      : David M. Rosenberg (dmr@davidrosenberg.me)
# @file        : Makefile
# @created     : Monday May 10, 2021 21:15:24 EDT
######################################################################

GO = go

OBJ = build/drchive

SRC = cmd/drchive/main.go \
			internal/db/db.go \
			internal/file/file.go \
			internal/log/log.go

TSTSRC = $(SRC) \
				 internal/db/db_test.go

build/drchive: $(SRC)
	$(GO) build -o build/drchive cmd/drchive/main.go

.PHONY: clean lint format cleanup

.DEFAULT: all

clean:
	rm -f $(OBJ)

format: $(SRC) $(TSTSRC)
	gofmt -s -w $(SRC) $(TSTSRC)

lint: $(SRC) $(TSTSRC)
	golint $(SRC) $(TSTSRC)

test:
	cd internal/db && $(GO) test -v

all: drchive
