grpc_go:
	protoc --go_out=youtube \
	--go_opt=paths=source_relative \
	--go-grpc_out=youtube \
	--go-grpc_opt=paths=source_relative --js_out=import_style=commonjs,binary:youtube --grpc-web_out=import_style=commonjs,mode=grpcwebtext:youtube youtube.proto
grpc_js:
	protoc \
  --proto_path=youtube \
  --js_out=import_style=commonjs:/youtube \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:/youtube \
  youtube.proto
