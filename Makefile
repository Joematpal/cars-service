# This is to keep out the 
bin/go.mod:
	@echo "// Hey, go mod, keep out!" > bin/go.mod

proto: go.mod bin/protoc bin/protobuf
	@./scripts/gen-proto-stubs

verify-proto: proto
	@./scripts/git-diff

bin/protoc: scripts/get-protoc
	@./scripts/get-protoc bin/protoc

bin/protobuf: bin/go.mod scripts/get-protoc-extras
	@./scripts/get-protoc-extras bin/protobuf


# bin/protoc-gen-go:
# 	# @go get -u google.golang.org/protobuf/protoc-gen-go