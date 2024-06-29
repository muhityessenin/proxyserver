FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

#dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#build
COPY cmd ./
RUN go build -o ./bin/cmd server/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/cmd /

EXPOSE 8080

CMD ["/cmd"]