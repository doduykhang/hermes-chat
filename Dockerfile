FROM golang:latest

RUN mkdir /build
WORKDIR /build
RUN cd /build

RUN export GO111MODULE=on
RUN git clone https://github.com/doduykhang/hermes-chat.git -b main .

RUN make clean
RUN make build

EXPOSE 8080

ENV LISTEN_PORT=""

ENTRYPOINT [ "build/build/app" ]

