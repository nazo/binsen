FROM alpine:3.8
ENV GOPATH /app
ENV PATH /app/bin:$PATH
RUN apk --no-cache --update add musl-dev go go-tools libtool postgresql-dev git bash make openssl readline autoconf automake curl bzip2 gzip tar zip gcc g++ && \
    rm -rf /var/cache/apk/*
WORKDIR /app/src/github.com/nazo/binsen/server
RUN mkdir -p $GOPATH/bin && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ADD ./Gopkg.toml /app/src/github.com/nazo/binsen/server
ADD ./Gopkg.lock /app/src/github.com/nazo/binsen/server
ADD ./Makefile /app/src/github.com/nazo/binsen/server
ADD ./bin /app/src/github.com/nazo/binsen/server/bin
RUN dep ensure -vendor-only && make deps-install
ADD . /app/src/github.com/nazo/binsen/server
