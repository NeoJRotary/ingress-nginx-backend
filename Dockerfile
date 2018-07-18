# build downloader
FROM golang:alpine as downloader
ENV GOPATH /build
WORKDIR /build
COPY ./downloader .
RUN apk --no-cache add git
RUN go get cloud.google.com/go/storage
RUN go get google.golang.org/api/iterator
RUN go build -o downloader .

# build reloader
FROM golang:alpine as reloader
ENV GOPATH /build
WORKDIR /build
COPY ./reloader .
RUN apk --no-cache add git
RUN go get github.com/NeoJRotary/exec-go
RUN go build -o reloader .

# build main nginx
FROM nginx:alpine
RUN apk --no-cache add bash ca-certificates curl
RUN rm /etc/nginx/conf.d/default.conf

# default ENV
ENV ENABLE_GCS_SYNC false
ENV GOOGLE_APPLICATION_CREDENTIALS /service_account.json
ENV CONFIGMAP_FOLDER /etc/config

COPY --from=downloader /build/downloader /downloader
COPY --from=reloader /build/reloader /reloader
COPY shells/. /
COPY module-conf /module-conf
CMD /start.sh
