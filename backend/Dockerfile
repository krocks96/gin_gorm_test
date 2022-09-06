FROM golang:1.18

COPY src/app /go/src/app

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apt-get update && apt-get install -y

RUN go mod tidy

CMD ["/bin/bash"]
