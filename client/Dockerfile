FROM node:14.9.0
WORKDIR /app
ADD ./package.json /app
ADD ./package-lock.json /app
RUN npm install
ADD . /app
