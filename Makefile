ifeq ($(OS), Windows_NT)
	RM = rmdir /s /q
else
	RM = rm
endif

create:
	mkdir gen
	protoc --proto_path=proto --go_out=gen/ ./proto/*.proto
	protoc --proto_path=proto --go_out=gen/ --go-grpc_out=gen/ ./proto/*.proto

clean:
	$(RM) .\gen\
