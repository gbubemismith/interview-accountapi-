FROM golang:latest

WORKDIR /app
COPY . .

RUN go get -d -v ./...

ENTRYPOINT [ "go", "test", "-v", "./accounts", "-coverprofile", "cover.out" ]