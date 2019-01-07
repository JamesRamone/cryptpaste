FROM golang:latest AS builder

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
RUN go build -o main . 

FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]