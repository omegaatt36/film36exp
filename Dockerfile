FROM golang:1.13.8-alpine3.11
WORKDIR /film36exp
ADD . /film36exp
RUN cd /film36exp && go build
EXPOSE 8080
ENTRYPOINT ./film36exp
