GOCMD=/usr/bin/go
GOBUILD=$(GOCMD) build

all: build

BINS=ex_fc_8_sqrt1 ex_fc_8_sqrt2 \
     ex_mtp_18_pic ex_mtp_23_wc ex_mtp_26_fib \
     ex_mtd_18_ipaddr ex_mtd_20_sqrt_err ex_mtd_22_reader ex_mtd_23_rot13 ex_mtd_25_pic2 \
	 ex_gen_2_list \
     ex_cnc_8_tree ex_cnc_10_crawler

build: $(BINS)

ex_fc_8_sqrt1: ex_fc_8_sqrt1.go
	$(GOBUILD) $<

ex_fc_8_sqrt2: ex_fc_8_sqrt2.go
	$(GOBUILD) $<

ex_mtp_18_pic: ex_mtp_18_pic.go tour/pic/pic.go
	$(GOBUILD) $<

ex_mtp_23_wc: ex_mtp_23_wc.go tour/wc/wc.go
	$(GOBUILD) $<

ex_mtp_26_fib: ex_mtp_26_fib.go
	$(GOBUILD) $<

ex_mtd_18_ipaddr: ex_mtd_18_ipaddr.go
	$(GOBUILD) $<

ex_mtd_20_sqrt_err: ex_mtd_20_sqrt_err.go
	$(GOBUILD) $<

ex_mtd_22_reader: ex_mtd_22_reader.go tour/reader/validate.go
	$(GOBUILD) $<

ex_mtd_23_rot13: ex_mtd_23_rot13.go
	$(GOBUILD) $<

ex_mtd_25_pic2: ex_mtd_25_pic2.go tour/pic/pic.go
	$(GOBUILD) $<

ex_gen_2_list: ex_gen_2_list.go
	$(GOBUILD) $<

ex_cnc_8_tree: ex_cnc_8_tree.go tour/tree/tree.go
	$(GOBUILD) $<

ex_cnc_10_crawler: ex_cnc_10_crawler.go
	$(GOBUILD) $<

clean:
	rm -rf $(BINS)

