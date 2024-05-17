// protobuf file wil be in proto dir 

cmd : protoc --go-grpc_out=. --go_out=. *.proto

// auto generates stubs : syntax_grpc.pb.go & syntax.pb.go

protoc          --> proto compiler
--go-grpc_out=. --> generate code containing grpc in current dir
--go_out=.      --> generate models in current dir
*.proto         --> (*) all proto files ; (.) current dir ; compile all

NOTE : for specific .proto --> replace *.proto with `file.proto`

// auto generates stubs : file_grpc.pb.go & file.pb.go


// if u get any path related errors while compiling .protorun below cmds 

# Add the bin directory to your PATH
export PATH="$HOME/.local/bin:$PATH"

# Next, install the Go Protocol Buffers plugin and the gRPC plugin:
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Ensure that your $GOPATH/bin is in your system's PATH
export PATH="$PATH:$(go env GOPATH)/bin"