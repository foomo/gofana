FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN adduser -D -u 1001 -g 1001 gofana

COPY gofana /usr/bin/

USER gofana
WORKDIR /home/gofana

ENTRYPOINT ["gofana"]
