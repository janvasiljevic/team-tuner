FROM golang:1.19-bullseye

WORKDIR /app 

ENV GOBIN=/go/bin
ENV PATH="${GOBIN}:${PATH}"

RUN go install github.com/cosmtrek/air@v1.26.0
RUN chmod +x /go/bin/air

RUN air -v

CMD ["air", "-c", ".air.toml"]