FROM nginx:alpine
ARG GO_SITE=web1
ARG GO_PORT=8000
ADD ${GO_SITE}.conf /etc/nginx/conf.d/default.conf
RUN adduser -D -H -u 1000 -s /bin/bash www-data