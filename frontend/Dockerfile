# The builder from node image
FROM node:11-alpine as builder

RUN mkdir /app
WORKDIR /app
COPY . .
FROM nginx:alpine
COPY ./.nginx/nginx.conf /etc/nginx/nginx.conf
RUN mkdir /etc/nginx/ssl/
COPY certs/donatuplasma.org.cer /etc/nginx/ssl/donatuplasma.org.cer
COPY certs/donatuplasma.org.key /etc/nginx/ssl/donatuplasma.org.key
COPY ./dist /usr/share/nginx/html
EXPOSE 80 80
EXPOSE 443 443

ENTRYPOINT ["nginx", "-g", "daemon off;"]