# build
FROM  golang:alpine  AS build-env

ADD  main.go /src/main.go
RUN  cd /src && go build -o /app/app main.go


# run
FROM alpine:latest

WORKDIR /app/
COPY --from=build-env /app/app /app/

EXPOSE 8080
ENTRYPOINT ./app

