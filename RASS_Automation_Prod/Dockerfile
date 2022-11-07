FROM golang:1.18 AS builder


WORKDIR /root

COPY . .

COPY *.go ./

#go env -w GO111MODULE=off
RUN go mod download -x 
RUN go build



FROM debian:buster-slim AS Final

WORKDIR /root

RUN apt-get update && apt install mariadb-client -y

COPY --from=builder /root/Rass .
COPY --from=Builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 5152

CMD ["./Rass"]

    

