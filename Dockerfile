FROM golang:1.22-alpine as build

WORKDIR /go/src/app

COPY ["go.mod", "go.sum", "./"]

RUN ["go", "mod", "download"]

COPY . .

ENV APP_NAME=film36exp

RUN ["go", "build", "-o", "build/${APP_NAME}" ,"./cmd/api"]

# FROM build as dev

# CMD ["go", "run", "./cmd/api"]

FROM gcr.io/distroless/static-debian12 as prod

WORKDIR /home/app/

COPY --from=build /go/src/app/build/${APP_NAME} ./

CMD ["./${APP_NAME}"]
