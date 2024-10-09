ARG ARCH=amd64
# build stage
FROM docker.io/golang:1.19 AS builder
RUN mkdir -p /go/src/rest2s3
WORKDIR /go/src/rest2s3
COPY . ./
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$ARCH go build -a -o /app .


# final stage
FROM docker.io/alpine:latest
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /app ./
COPY --from=builder /go/src/rest2s3/idx.tmpl ./
RUN chmod +x ./app
ENTRYPOINT ["./app"]
EXPOSE 8080