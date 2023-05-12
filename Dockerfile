FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ard2rss .


FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /app/ard2rss /app/ard2rss
EXPOSE 8080
USER nonroot:nonroot

CMD ["/app/ard2rss"]
