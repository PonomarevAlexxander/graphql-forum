FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /graphql-forum ./cmd/app/main.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /graphql-forum /graphql-forum

EXPOSE 8080

ENTRYPOINT ["/graphql-forum"]
