FROM golang:alpine AS builder

WORKDIR /app

ADD . ./
RUN go mod download
RUN go build -buildvcs=false -o /app/schedule-manager-backend ./cmd/app


FROM alpine:latest
USER goapp
WORKDIR /app
COPY --from=builder /app/schedule-manager-backend  /app/
COPY --from=builder /app/configs /app/configs/
COPY --from=builder /app/templates /app/templates/

CMD ["./schedule-manager-backend"]
