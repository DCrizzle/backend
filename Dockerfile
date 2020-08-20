FROM golang:1.13

ENV GO111MODULE=on \
		CGO_ENABLED=0 \
    GOOS=linux \
		GOARCH=amd64

WORKDIR /build

COPY helper/ .

RUN go mod init helper

RUN go mod download > /dev/null

RUN go build -o helper .

WORKDIR /app

RUN cp /build/helper .
RUN cp /build/config.json .

EXPOSE 4080

CMD ["./helper"]
