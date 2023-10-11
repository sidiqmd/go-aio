# Stage 1
FROM node:16-alpine AS node

WORKDIR /app

COPY frontend/package*.json ./
RUN npm install

COPY frontend/. .

RUN npm run build --prod

# Stage 2
FROM nginx:alpine

COPY --from=node /app/dist/frontend /usr/share/nginx/html

EXPOSE 80