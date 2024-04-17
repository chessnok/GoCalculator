FROM node:21.5-alpine3.19 as build
COPY ./frontend /app
WORKDIR /app
RUN npm install
RUN npm run build
FROM nginx:1.25-alpine as prod
COPY --from=build /app/build /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]