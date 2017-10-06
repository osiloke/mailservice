FROM alpine:latest
LABEL Osiloke Emoekpere <me@osiloke.com>
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD mailsender /app/mailsender 
WORKDIR /app
CMD ["/app/mailsender", "machinery"]