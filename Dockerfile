FROM golang:1.13

ENV GO111MODULE=on \
		CGO_ENABLED=0 \
    GOOS=linux \
		GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod download > /dev/null

# compile backend binary
RUN go build -o backend .

WORKDIR /app

RUN cp /build/backend .
RUN cp /build/start .
RUN cp /build/database/schema.graphql .

RUN apt-get update
RUN apt-get install -y curl

# download dgraph binary
RUN curl https://get.dgraph.io -sSf | bash && dgraph

EXPOSE 8888

CMD ["bash", "start"]

# outline:
# [ ] set "general" ubuntu image
# [ ] run install from source steps
# - [ ] https://github.com/dgraph-io/dgraph#install-from-source
# [ ] expose required port
# [ ] run bash start script
