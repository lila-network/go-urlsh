FROM golang:1.20-alpine AS dev

COPY . /var/app
WORKDIR /var/app

ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux

EXPOSE 80
ENTRYPOINT ["sh"]


FROM dev as build

RUN go mod download && go mod verify
RUN go build -o go-urlsh main.go
RUN chmod +x go-urlsh


FROM scratch AS prod

WORKDIR /var/app
COPY --from=build /var/app/go-urlsh /var/app/
COPY --from=build /var/app/web /var/app/web

EXPOSE 80
ENTRYPOINT ["/var/app/go-urlsh"]
CMD ""