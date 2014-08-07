GOCMD=/usr/bin/go
GOBUILD=$(GOCMD) build

all: build

BINS=ex25_sqrt1 ex25_sqrt2 ex38_pic ex43_wc ex46_fib ex50_cbrt ex58_sqrt_err ex59_http ex60_http_hdlr ex62_pic2 ex63_rot13

build: $(BINS)

ex25_sqrt1: ex25_sqrt1.go
	$(GOBUILD) $<

ex25_sqrt2: ex25_sqrt2.go
	$(GOBUILD) $<

ex38_pic: ex38_pic.go
	$(GOBUILD) $<

ex43_wc: ex43_wc.go
	$(GOBUILD) $<

ex46_fib: ex46_fib.go
	$(GOBUILD) $<

ex50_cbrt: ex50_cbrt.go
	$(GOBUILD) $<

ex58_sqrt_err: ex58_sqrt_err.go
	$(GOBUILD) $<

ex59_http: ex59_http.go
	$(GOBUILD) $<

ex60_http_hdlr: ex60_http_hdlr.go
	$(GOBUILD) $<

ex62_pic2: ex62_pic2.go
	$(GOBUILD) $<

ex63_rot13: ex63_rot13.go
	$(GOBUILD) $<

clean:
	rm -rf $(BINS)

