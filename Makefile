IPLOC_CONV = $(GOPATH)/bin/iploc-conv
IPLOC_FETCH := $(GOPATH)/bin/iploc-fetch
IPLOC_GEN = $(GOPATH)/bin/iploc-gen


.PHONY: all test clean install prepare

all: install prepare test

clean:
		rm -f $(IPLOC_CONV) $(IPLOC_FETCH) $(IPLOC_GEN)
		rm -f qqwry.dat

install:
		cd cmd/iploc-conv; go install
		cd cmd/iploc-fetch; go install
		cd cmd/iploc-gen; go install

prepare:
		iploc-fetch ./qqwry.gbk.dat -q
		iploc-conv -s qqwry.gbk.dat -d qqwry.dat -n -q
		rm qqwry.gbk.dat

test:
		go test -v

bench:
		go test -bench=.