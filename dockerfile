FROM golang:1.20-buster

WORKDIR /app
COPY go.sum go.mod /app/

RUN go mod download
COPY . /app

RUN CGO_ENABLED=0 go build -o /bin/app ./cmd/

FROM scratch
COPY --from=0 /bin/app /bin/app

ENTRYPOINT ["/bin/app"]