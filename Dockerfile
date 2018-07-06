# Se baseia na imagem https://hub.docker.com/_/golang/
FROM golang:latest

# Instala godep
RUN go get -u github.com/golang/dep/cmd/dep

# Copia o diretorio local para o diretorio do container
ADD . $GOPATH/src/github.com/michelaquino/golang_api_skeleton

# Seta diretório de trabalho
WORKDIR /go/src/github.com/michelaquino/golang_api_skeleton

# Instala dependencias
RUN make setup

# Compila aplicação
RUN GOOS=linux GOARCH=amd64 go build -o golang_api_skeleton main.go

# Executa a aplicacao quando o container for iniciado
ENTRYPOINT /go/src/github.com/michelaquino/golang_api_skeleton/golang_api_skeleton

# Expoe a porta 8080
EXPOSE 8080