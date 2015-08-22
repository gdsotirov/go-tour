GOCMD=/usr/bin/go
GOBUILD=$(GOCMD) build

all: build

BINS=ex_fc_8_sqrt1 ex_fc_8_sqrt2 ex_mtp_14_pic ex_mtp_19_wc ex_mtp_22_fib ex_mtd_7_ipaddr ex_mtd_9_sqrt_err ex_mtd_12_rot13 ex_mtd_14_http_hdlr ex_mtd_16_pic2 ex_cnc_8_tree ex_cnc_10_crawler

build: $(BINS)

ex_fc_8_sqrt1: ex_fc_8_sqrt1.go
	$(GOBUILD) $<

ex_fc_8_sqrt2: ex_fc_8_sqrt2.go
	$(GOBUILD) $<

ex_mtp_14_pic: ex_mtp_14_pic.go
	$(GOBUILD) $<

ex_mtp_19_wc: ex_mtp_19_wc.go
	$(GOBUILD) $<

ex_mtp_22_fib: ex_mtp_22_fib.go
	$(GOBUILD) $<

ex_mtd_7_ipaddr: ex_mtd_7_ipaddr.go
	$(GOBUILD) $<

ex_mtd_9_sqrt_err: ex_mtd_9_sqrt_err.go
	$(GOBUILD) $<

ex_mtd_12_rot13: ex_mtd_12_rot13.go
	$(GOBUILD) $<

ex_mtd_14_http_hdlr: ex_mtd_14_http_hdlr.go
	$(GOBUILD) $<

ex_mtd_16_pic2: ex_mtd_16_pic2.go
	$(GOBUILD) $<

ex_cnc_8_tree: ex_cnc_8_tree.go
	$(GOBUILD) $<

ex_cnc_10_crawler: ex_cnc_10_crawler.go
	$(GOBUILD) $<

clean:
	rm -rf $(BINS)

