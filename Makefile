ifeq ($(OS), Windows_NT)
	RM = rmdir /s /q
else
	RM = rm
endif

create:
	mkdir gen
# mkdir services\youtube-dl\gen
	protoc --proto_path=proto --go_out=gen/ ./proto/*.proto
	protoc --proto_path=proto --go_out=gen/ --go-grpc_out=gen/ ./proto/*.proto
#powershell conda activate grpc && python -m grpc_tools.protoc --proto_path=proto --python_out=services/youtube-dl/gen/ --grpc_python_out=services/youtube-dl/gen/ ./proto/*.proto

clean:
	$(RM) .\gen\
