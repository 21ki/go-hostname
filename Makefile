PROGRAM       := go-hostname
.DEFAULT_GOAL := $(PROGRAM)
AWK           := awk
RM            := rm
GO            := go
GOFMT         := gofmt

$(PROGRAM): $(PROGRAM).go
	$(GO) build $<

.PHONY: lint
lint: $(PROGRAM).go
	$(GOFMT) -w $<

.PHONY: clean
clean:
	@if [ -f $(PROGRAM) ]; then \
	  $(RM) -fv $(PROGRAM); \
	fi
