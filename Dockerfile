FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go build -o /go-backend
CMD ["/go-backend"]

FROM nginx:latest
RUN apt-get update && \
    apt-get install -y certbot python-certbot-nginx
COPY nginx.conf /etc/nginx/nginx.conf
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
EXPOSE 80
EXPOSE 443
ENTRYPOINT ["/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]