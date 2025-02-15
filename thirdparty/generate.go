package thirdparty

//go:generate mkdir -p github.com/envoyproxy/protoc-gen-validate/validate
//go:generate curl -fsSL https://raw.githubusercontent.com/bufbuild/protoc-gen-validate/main/validate/validate.proto -o github.com/envoyproxy/protoc-gen-validate/validate/validate.proto

//go:generate mkdir -p github.com/mwitkow/go-proto-validators
//go:generate curl -fsSL https://github.com/mwitkow/go-proto-validators/raw/master/validator.proto -o github.com/mwitkow/go-proto-validators/validator.proto

//go:generate mkdir -p github.com/pseudomuto/protokit/fixtures
//go:generate curl -fsSL https://github.com/pseudomuto/protokit/raw/master/fixtures/extend.proto -o github.com/pseudomuto/protokit/fixtures/extend.proto

//go:generate mkdir -p github.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/
//go:generate curl -fsSL https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/annotations.proto -o github.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/annotations.proto
//go:generate curl -fsSL https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/openapiv2.proto -o github.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/openapiv2.proto
