FROM golang:1.21-alpine
WORKDIR /home/app
COPY . .
RUN go build -o raceApi
CMD ./raceApi