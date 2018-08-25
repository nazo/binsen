FROM alpine:3.8
ENV GOPATH /app
ENV PATH /app/bin:$PATH
RUN apk --no-cache --update add musl-dev go go-tools libtool postgresql-dev git bash make openssl readline autoconf automake curl bzip2 gzip tar zip gcc g++ && \
    go get -u github.com/golang/dep/cmd/dep && \
    go get -u gopkg.in/godo.v2/cmd/godo && \
    go get -u github.com/pressly/goose/cmd/goose && \
    go get -u github.com/volatiletech/sqlboiler && \
    go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql && \
    rm -rf /var/cache/apk/*
WORKDIR /app/src/github.com/nazo/binsen/server
ADD ./Gopkg.toml /app/src/github.com/nazo/binsen/server
ADD ./Gopkg.lock /app/src/github.com/nazo/binsen/server
RUN dep ensure -vendor-only
ADD . /app/src/github.com/nazo/binsen/server
