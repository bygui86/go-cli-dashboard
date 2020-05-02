
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## example-app

build-example-app :		## Build example application
	cd ./example-app && \
	make build

run-example-app :		## Run example application
	cd ./example-app && \
	make run

## cli-dashboard

### termui based

build-termui-based :		## Build CLI dashboard based on gizak/termui
	cd ./termui-based && \
	make build

run-termui-based :		## Run CLI dashboard based on gizak/termui
	cd ./termui-based && \
	make run TARGET_HOST=localhost TARGET_PORT=6060

### termbash based

build-termdash-based :		## Build CLI dashboard based on mum4k/termdash
	cd ./termdash-based && \
	make build

run-termdash-based :		## Run CLI dashboard based on mum4k/termdash
	cd ./termdash-based && \
	make run

## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
