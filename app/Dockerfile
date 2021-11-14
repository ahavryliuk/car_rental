FROM golang:1.16-alpine

ENV \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/cr

COPY go.* ./
RUN go mod download

COPY . .

RUN cd ./cmd/rental_api; go build -gcflags "all=-N -l" -o /cr-bin/rental_api -x
RUN cd ./cmd/bookings_logger; go build -gcflags "all=-N -l" -o /cr-bin/bookings_logger -x

COPY config/ /cr-bin/config

EXPOSE 2000

CMD ["/bin/sh", "run.sh"]