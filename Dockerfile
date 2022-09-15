FROM golang:1.18 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /build

COPY switchboard-server/ .
RUN go mod download && \
    go build -o main

FROM gcr.io/distroless/static-debian11:latest
WORKDIR /
COPY --from=builder /build/main /main
USER nonroot
CMD [ "/main" ]
