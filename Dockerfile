FROM golang:1.22-alpine as build

WORKDIR /go/src/app

COPY ["go.mod", "go.sum", "./"]

RUN ["go", "mod", "download"]

COPY . .

ARG CMD

# 使用 shell 形式來執行指令
RUN go build -o build/app ./cmd/${CMD}

# FROM build as dev

# CMD ["go", "run", "./cmd/api"]

FROM gcr.io/distroless/static-debian12 as prod

WORKDIR /home/app/

COPY --from=build /go/src/app/build/app ./

CMD ["./app"]