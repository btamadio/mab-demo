FROM golang:1.16 AS build

RUN mkdir /reward
WORKDIR /reward
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/reward

FROM scratch
COPY --from=build /go/bin/reward /go/bin/reward
EXPOSE 1337

ENTRYPOINT ["/go/bin/reward"]
