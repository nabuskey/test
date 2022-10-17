# Base Image
FROM golang:1.19.2 as builder

ENV GOPROXY=direct
#working directory
RUN mkdir -p /app
WORKDIR /app

# Dependecy installation
ADD go.mod go.mod
ADD go.sum go.sum
# RUN go mod download -json
# Copy code
ADD . .
# build the source
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Final Image
FROM alpine:3.16
# add ca-certificates
RUN apk update && apk --no-cache  add ca-certificates
# set working directory
WORKDIR /app
# copy the binary from builder
COPY --from=builder /app/main ./main
# run the binary
CMD ["./main"]
