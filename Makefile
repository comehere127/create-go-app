# A default Makefile for Create Go App project.
# Author: Vic Shóstak <koddr.me@gmail.com> (https://github.com/koddr)
# For more information, please visit https://github.com/comehere127/create-go-app

.PHONY: test run build

BACKEND_PATH = $(PWD)/backend

test:
	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && go test ./...; fi

run: test
	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && $(MAKE) run; fi

build: test
	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && $(MAKE) build; fi
