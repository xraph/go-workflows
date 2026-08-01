# Defines `fmt` and `test` only. `clippy`, `audit`, `build` and `docs` are
# deliberately absent so rust-ci.yml must fall back to cargo for those.
.PHONY: fmt test

fmt:
	@echo "MAKE-TARGET-FMT"
	cargo fmt --all -- --check

test:
	@echo "MAKE-TARGET-TEST"
	cargo test
