FROM golang:1.25 AS build
WORKDIR /app
COPY . .
RUN bash ./scripts/build-application.sh

FROM alpine
COPY --from=build /app/ordersystem /app/ordersystem
CMD ["/app/ordersystem"]