init:
	go mod tidy
	make generate_mocks

clean:
	find . -name "*.mock.gen.go" -type f -delete

# Generate mock files
GO_FILES := $(shell find ./ -name "interfaces.go")
GEN_GO_FILES := $(GO_FILES:%.go=%.mock.gen.go)

generate_mocks: $(GEN_GO_FILES)
$(GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))

test_unit:
	go test -short -coverprofile=coverage.out -v ./... -coverpkg=./...