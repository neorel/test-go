FROM golang:1.21-alpine
WORKDIR /home/go

# CMD ["sleep", "1d"]

COPY . .
RUN go build -o raceApi
CMD ./raceApi