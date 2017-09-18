FROM golang:1.9

RUN mkdir -p /app
WORKDIR /app
ADD . /app

CMD ["go", "run", "main.go"]
