FROM golang:alpine as buildApp
ENV GOPATH /build
WORKDIR /build
COPY . .
RUN apk --no-cache add git
RUN go get cloud.google.com/go/storage
RUN go get google.golang.org/api/iterator
RUN go build -o initConfig .

FROM nginx:alpine
RUN apk --no-cache add bash ca-certificates curl
RUN rm /etc/nginx/conf.d/default.conf

ENV GOOGLE_APPLICATION_CREDENTIALS /service_account.json

COPY --from=buildApp /build/initConfig /initConfig
COPY start.sh /start.sh
CMD /start.sh
