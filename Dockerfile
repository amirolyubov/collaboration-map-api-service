FROM golang:1.20 AS builder

RUN go version
RUN apt-get install git

COPY ./ /cm-api
WORKDIR /cm-api

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 go build -o ./app

# second image from first one, but without preinstalled golang 
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /cm-api/app .
EXPOSE 8080

CMD [ "./app" ]
