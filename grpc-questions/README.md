
# WIP

Para gerar o stub do lado do servidor, como está escrito em golang, basta usar o comando abaixo:

```console
$ protoc --go_out=. --go_opt=paths=source_relative \
     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
     questions/questions.proto
```

Para gerar o stub em dart basta usar o comando abaixo:

```
$ protoc --dart_out=grpc:questions_app/lib/model questions/questions.proto
```

Artigos usados para escrever o aplicativo em flutter 
- [Flutter ❤ GRPC](https://medium.com/flutter-community/flutter-grpc-810f87612c6d)