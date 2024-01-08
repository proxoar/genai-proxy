FROM golang:1.21-alpine AS builder

RUN apk update && apk add --no-cache make

COPY ${PWD} /app
WORKDIR /app
RUN make build

FROM alpine

RUN apk --update add ca-certificates && \
    rm -rf /var/cache/apk/*

RUN adduser -D appuser
USER appuser

WORKDIR /home/appuser/app

COPY --from=builder /app/main /home/appuser/app/appbin

EXPOSE 8000

CMD ["./appbin"]