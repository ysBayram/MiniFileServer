FROM golang:1.14

WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch
COPY --from=0 /bin/app /bin/app
ENTRYPOINT ["/bin/app"]