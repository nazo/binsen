FROM node:14.13.1
WORKDIR /app
ADD ./package.json /app
ADD ./package-lock.json /app
RUN npm install
ADD . /app
