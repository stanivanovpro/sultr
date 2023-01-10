FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json .
RUN yarn global add @quasar/cli
COPY . .
RUN yarn
RUN quasar build

FROM nginx:stable-alpine as publish-stage
COPY --from=build-stage /app/dist/spa /usr/share/nginx/html
COPY --from=build-stage /app/deployment/default.conf /etc/nginx/conf.d/default.conf
CMD ["nginx", "-g", "daemon off;"]
EXPOSE 8080
