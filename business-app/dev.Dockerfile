FROM node:lts-alpine as install-stage
WORKDIR /app
COPY package*.json .
RUN yarn global add @quasar/cli
COPY . .
RUN yarn

FROM install-stage as publish-stage
WORKDIR /app
CMD quasar dev --port 4000
EXPOSE 4000
