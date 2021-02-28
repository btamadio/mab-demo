FROM golang:1.16 AS build

RUN mkdir /bandit
WORKDIR /bandit
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/bandit

FROM scratch
COPY --from=build /go/bin/bandit /go/bin/bandit
EXPOSE 1338

ENTRYPOINT ["/go/bin/bandit"]
