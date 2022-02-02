FROM golang:1.16.3

WORKDIR /app

COPY go.mod .
COPY go.sum . 

RUN go mod download 

COPY . /app

EXPOSE 8282

CMD ["make", "all"]