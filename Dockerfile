FROM golang:1.13

ENV GO111MODULE=on \
		CGO_ENABLED=0 \
    GOOS=linux \
		GOARCH=amd64

WORKDIR /backend

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["./bin/start_backend"]
