FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.io,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/otp api/otp.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/otp /app/otp
COPY --from=builder /app/etc /app/etc

CMD ["./otp", "-f", "etc/otp.yaml"]
