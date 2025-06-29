FROM golang:latest as build
COPY .. /app/
WORKDIR /app/
RUN CGO_ENABLED=0 go build -o main.go

FROM alpine as prod
COPY --from=build /app/main.go /app/
COPY ./config.yaml /app/

ENTRYPOINT [ "/app/main.go" ]