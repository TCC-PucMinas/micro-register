# micro-register


## Generate
```sh
$ protoc -I communicate/ communicate/authenticate.proto --go_out=plugins=grpc:communicate
$ protoc -I communicate/ communicate/users.proto --go_out=plugins=grpc:communicate
```


protoc -I=communicate --go_out=. --go-grpc_out=. communicate/.proto protoc --go_out=. --proto_path=./communicate communicate/.proto --go-grpc_out=./communicate --go-grpc_opt=paths=source_relative communicate/slack.proto



// DY479456577BR