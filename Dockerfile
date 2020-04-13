FROM golang AS build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

FROM alpine

COPY --from=build /server /server

ENTRYPOINT ["/server"]
