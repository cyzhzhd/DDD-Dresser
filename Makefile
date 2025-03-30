PROJECT_NAME ?= earth
PLUGIN_OUTPUT_DIR ?= $(CURDIR)/dist

run:
	go run main.go
running:
	export PATH=${PATH}:`go env GOPATH`/bin && \
	reflex -r '\.go$$' \
		-s -- sh -c 'go run main.go'
