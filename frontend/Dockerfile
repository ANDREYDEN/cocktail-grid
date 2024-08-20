FROM node:20

WORKDIR /frontend

COPY package.json ./
COPY yarn.lock ./
RUN yarn

COPY . .