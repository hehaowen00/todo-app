# reverse proxy

FROM nginx:stable-alpine

COPY ./certs /etc/nginx/certs

COPY ./nginx/nginx.conf /etc/nginx/conf.d/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]