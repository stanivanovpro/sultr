# build stage
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json .
RUN yarn global add @quasar/cli
COPY . .
RUN yarn
RUN quasar build

FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist/spa /usr/share/nginx/html
COPY --from=build-stage /app/deployment/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
