# stage 0

FROM node:16.15-alpine

WORKDIR /app

COPY package.json ./

COPY package-lock.json ./

RUN npm install --silent

COPY . .

RUN npm run build

# stage 1

FROM nginx:stable-alpine

COPY --from=0 /app/build /usr/share/nginx/html

COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]