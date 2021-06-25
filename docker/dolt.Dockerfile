FROM alpine:latest

RUN apk add --update-cache \
    bash curl mysql-client \
  && rm -rf /var/cache/apk/*

RUN bash -c 'curl -L https://github.com/dolthub/dolt/releases/latest/download/install.sh | bash'

WORKDIR /var/lib/dolt

RUN adduser -D -u 1000 dolt
USER dolt

CMD dolt sql-server 
