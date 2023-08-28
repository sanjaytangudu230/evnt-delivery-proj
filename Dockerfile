FROM golang:1.21

WORKDIR /app

COPY . .

RUN make build

EXPOSE 8080

CMD ["make", "start"]