FROM golang:1.23 as builder
WORKDIR /app
COPY . .
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM scratch
COPY --from=builder /app/cmd/server /server
ENTRYPOINT [ "/server" ]
