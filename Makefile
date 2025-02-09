all: dev

.PHONY: go-gen
go-gen:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	cd server && buf generate

.PHONY: ts-gen
ts-gen:
	cd site && npx buf generate

.PHONY: gen
gen: go-gen ts-gen

.PHONY: run-site
run-site:
	cd site && npm install && npm run dev

.PHONY: run-server
run-server:
	go install github.com/mitranim/gow@latest
	cd server && gow run ./... -c ../settings.toml

.PHONY: dev
dev: gen
	make -j 2 run-site run-server
