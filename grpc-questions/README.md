
# WIP

Para gerar o stub do lado do servidor, como está escrito em golang, basta usar o comando abaixo:

```console
$ protoc --go_out=. --go_opt=paths=source_relative \
     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
     questions/questions.proto
```