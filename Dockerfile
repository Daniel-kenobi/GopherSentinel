FROM golang
LABEL authors="daniel-kenobi"

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify


COPY . .
RUN go build

CMD ./GopherSentinel