FROM golang:1.13

ENV GO111MODULE=on \
		CGO_ENABLED=0 \
    GOOS=linux \
		GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod download > /dev/null

# compile helper binary
RUN go build -o helper .

WORKDIR /app

RUN cp /build/helper .
RUN cp /build/start .
RUN cp /build/database/schema.graphql .

RUN apt-get update
RUN apt-get install -y curl

# download dgraph binary
# outline:
# [ ] run install from source steps (for master branch access)
# - [ ] https://github.com/dgraph-io/dgraph#install-from-source
RUN curl https://get.dgraph.io -sSf | bash && dgraph

EXPOSE 8888

CMD ["bash", "start"]
