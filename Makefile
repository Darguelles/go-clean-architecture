gen:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate types --package openapi -o internal/controller/openapi/types.gen.go openapi.yaml
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate spec --package openapi -o internal/controller/openapi/spec.gen.go openapi.yaml
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate server --package openapi -o internal/controller/openapi/server.gen.go openapi.yaml

start:
	go run .
