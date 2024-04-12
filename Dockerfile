FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go build -o /go-backend
CMD ["/go-backend"]

FROM nginx:latest
RUN apt-get update && \
    apt-get install -y wget gnupg2 software-properties-common && \
    wget -q https://dl.eff.org/certbot-auto && \
    mv certbot-auto /usr/local/bin/certbot-auto && \
    chown root /usr/local/bin/certbot-auto && \
    chmod 0755 /usr/local/bin/certbot-auto && \
    /usr/local/bin/certbot-auto --non-interactive --install-only
RUN /usr/local/bin/certbot-auto --non-interactive plugins && \
    apt-get install -y python3-certbot-nginx
COPY nginx.conf /etc/nginx/nginx.conf
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
EXPOSE 80
EXPOSE 443
ENTRYPOINT ["/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
