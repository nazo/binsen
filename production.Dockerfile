FROM alpine:3.8
ENV GOPATH /server
ENV PATH /server/bin:/client/node_modules/.bin:$PATH
RUN apk --no-cache --update add musl-dev go go-tools libtool postgresql-dev git bash make openssl readline autoconf automake curl bzip2 gzip tar zip gcc g++ nodejs nodejs-dev nodejs-npm curl-dev bzip2-dev libressl-dev ncurses nginx && \
    mkdir -p /client && \
    mkdir -p /server && \
    go get -u github.com/golang/dep/cmd/dep && \
    go get -u gopkg.in/godo.v2/cmd/godo && \
    go get -u github.com/pressly/goose/cmd/goose && \
    go get -u github.com/volatiletech/sqlboiler && \
    go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql && \
    rm -rf /var/cache/apk/*

WORKDIR /client
ADD ./client/package.json /client
ADD ./client/package-lock.json /client
RUN npm install

WORKDIR /server/src/github.com/nazo/binsen/server
ADD ./server/Gopkg.toml /server/src/github.com/nazo/binsen/server
ADD ./server/Gopkg.lock /server/src/github.com/nazo/binsen/server
RUN dep ensure -vendor-only

ADD ./.circleci /deploy

ADD ./client /client
RUN cd /client && npm run build

ADD ./server /server/src/github.com/nazo/binsen/server
RUN cd /server/src/github.com/nazo/binsen/server && go build -o /server/binsen cmd/main.go
