# micro-register


## Generate
```sh
$ protoc --go_out=paths=source_relative:../communicate -I. authenticate.proto
```


protoc -I=communicate --go_out=. --go-grpc_out=. communicate/.proto protoc --go_out=. --proto_path=./communicate communicate/.proto --go-grpc_out=./communicate --go-grpc_opt=paths=source_relative communicate/slack.proto