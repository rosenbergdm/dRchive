######################################################################
# @author      : David M. Rosenberg (dmr@davidrosenberg.me)
# @file        : Makefile
# @created     : Monday May 10, 2021 21:15:24 EDT
######################################################################

GO = go

OBJ = build/drchive

build/drchive: cmd/drchive/main.go internal/db/db.go
	$(GO) build -o build/drchive cmd/drchive/main.go

drchive: $(OBJ)
	$(GO) build -o build/drchive cmd/drchive/main.go

.PHONY: clean all

.DEFAULT: all

clean:
	rm -f $(OBJ)

test:
	cd include/db && $(GO) test

all: drchive
