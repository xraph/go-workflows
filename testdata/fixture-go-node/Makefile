# This target FAILS on purpose. The smoke job calls the workflow with
# prefer-makefile: false, so a passing run proves the Makefile was bypassed.
# If the probe ever regresses to preferring make, this fixture goes red.
.PHONY: test

test:
	@echo "MAKEFILE-TEST-SHOULD-NOT-RUN"
	@exit 1
