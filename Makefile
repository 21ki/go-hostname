PROGRAM       := go-hostname
.DEFAULT_GOAL := $(PROGRAM)
AWK           := awk

$(PROGRAM): $(PROGRAM).go
	go build $@.go

clean:
	@if [ -f $(PROGRAM) ]; then \
	  rm -fv $(PROGRAM); \
	fi
