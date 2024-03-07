FROM golang:1.22.0

WORKDIR /app

COPY go.mod .
COPY . /test

RUN go build /test

ENTRYPOINT [ "./golang_testing_grpc" ]