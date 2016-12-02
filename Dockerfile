# Se baseis na imagem https://hub.docker.com/_/golang/
FROM golang:1.7.4

# Copia o diretorio local para o diretorio do container
ADD . $GOPATH/src/github.com/michelaquino/golang_API_skeleton

# Instala a aplicacao
RUN go install github.com/michelaquino/golang_API_skeleton

# Executa a aplicacao quando o container for iniciado
ENTRYPOINT $GOPATH/bin/golang_API_skeleton

# Expoe a porta 8080
EXPOSE 8080