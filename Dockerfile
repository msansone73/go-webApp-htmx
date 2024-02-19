FROM golang:1.21

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/server/...

CMD ["app"]

# gerar a imagem e fazer push
# docker build -t msansone/go-webapp-htmx .
# docker push msansone/go-webapp-htmx    