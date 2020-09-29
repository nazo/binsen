FROM golang:1.15.1-buster
ENV GOPATH /app
ENV PATH /app/bin:$PATH
ENV GO111MODULE on
WORKDIR /app/src/github.com/nazo/binsen/server
RUN go get -u gopkg.in/godo.v2/cmd/godo
ADD ./go.mod /app/src/github.com/nazo/binsen/server
ADD ./go.sum /app/src/github.com/nazo/binsen/server
ADD ./Makefile /app/src/github.com/nazo/binsen/server
ADD ./bin /app/src/github.com/nazo/binsen/server/bin
RUN go mod download && make tools-install
ADD . /app/src/github.com/nazo/binsen/server
