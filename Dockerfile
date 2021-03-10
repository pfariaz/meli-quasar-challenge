FROM golang AS build

WORKDIR /
COPY . .

RUN go build -o /quasar .

FROM golang AS bin

RUN mkdir /app
WORKDIR /app

COPY --from=build /quasar /app/quasar
COPY --from=build /satellites.json /app/satellites.json

CMD ["/app/quasar"]
