FROM golang:1.18-alpine as build
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app ./cmd/main.go


FROM gcr.io/distroless/base-debian11
WORKDIR /

EXPOSE 8080

COPY --from=build /go/bin/app /
USER nonroot:nonroot

CMD ["/app"]


