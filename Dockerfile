FROM golang:latest
WORKDIR /app
ENV Port=:1000
COPY . .
RUN go mod tidy && \
    go build -o myGolangApp
ENTRYPOINT ./myGolangApp