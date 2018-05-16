IPLOC := $(GOPATH)/bin/iploc
IPLOC_GEN = $(GOPATH)/bin/iploc-gen
IPLOC_CONV = $(GOPATH)/bin/iploc-conv

.PHONY: all test clean install

all: install test

clean:
		rm -f $(IPLOC) $(IPLOC_GEN) $(IPLOC_CONV)

install:
		cd cmd/iploc-gen; go install
		cd cmd/iploc; $(IPLOC_GEN) ../../qqwry.dat -n; go install
		cd cmd/iploc-conv; go install

test:
		go test -v

bench:
		go test -bench=.