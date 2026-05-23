# Tasks
- create the ansible playbook which perform following tasks
- installs 
    - latest release of go and update PATH, GOROOT, GOPATH in bashrc
    - latest release of protobuf-compiler

- verifies the above installation

- runs following to two command
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest