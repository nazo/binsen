FROM alpine:3.8
RUN apk --no-cache --update add nodejs nodejs-dev nodejs-npm curl-dev bzip2-dev libressl-dev git bash make readline autoconf automake curl bzip2 gzip tar zip gcc g++ ncurses && \
    rm -rf /var/cache/apk/*
WORKDIR /app
ADD ./package.json /app
ADD ./package-lock.json /app
RUN npm install
ADD . /app
