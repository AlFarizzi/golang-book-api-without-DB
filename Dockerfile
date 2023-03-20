FROM golang
ENV PORT 8000
EXPOSE 8000

WORKDIR /go/src/app
COPY . .

RUN go mod vendor
RUN go build -v -o app
RUN mv ./app /go/bin/

CMD ["app"]