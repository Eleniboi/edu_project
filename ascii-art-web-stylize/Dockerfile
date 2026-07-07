# stage 1
FROM golang:1.22-alpine AS export


WORKDIR /app


COPY go.mod ./

COPY . .


RUN go mod download 

RUN go build -o server .


# stage 2 

FROM alpine:3.19

WORKDIR /app

COPY --from=export /app/banners ./banners

COPY --from=export /app/template ./template

COPY --from=export /app/static ./static

EXPOSE 8000

CMD ["./server"]