FROM golang

EXPOSE 8080

WORKDIR /app
COPY . .
RUN go test ./...
RUN go build

ENV PORT 8080
ENV GIN_MODE=release

CMD ["/app/siribus"]
