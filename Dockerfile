FROM golang:1.14.0-buster
WORKDIR /film36exp
ADD . /film36exp
RUN cd /film36exp && go build
EXPOSE 8080
ENTRYPOINT ./film36exp
