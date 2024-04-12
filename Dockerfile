FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go build -o /go-backend
CMD ["/go-backend"]

FROM nginx:alpine
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d
