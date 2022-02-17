FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

WORKDIR /app/
ADD ./ /app/
# ADD ./zoneinfo.zip /usr/local/go/lib/time/
ADD ./demo.html /app/
ENTRYPOINT ["./app"]
