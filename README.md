go install github.com/golang/protobuf/proto@latest
go install github.com/golang/protobuf/protoc-gen-go@latest

After creation of chat file generate file for protoc
protoc --go_out=. chat.proto

![image](https://github.com/Mohan-Website/gRPC/assets/77502414/c4bb590f-4cc0-4cf9-a529-e8b882cc3278)
