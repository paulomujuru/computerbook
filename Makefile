gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:protobuf

clean:
	rm protobuf/*.go

run:
	go run main.go