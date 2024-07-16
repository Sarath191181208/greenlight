MAIN_PACKAGE_PATH := ./cmd/api
BINARY_NAME := api

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	air \
		--build.cmd "go build -o /tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}" \
		--build.bin "./bin/api" \
		--build.exclude_dir "requests"
