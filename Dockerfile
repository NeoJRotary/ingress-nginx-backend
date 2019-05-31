# build downloader
FROM golang:stretch as downloader
ENV GOPATH /build
WORKDIR /build
RUN go get cloud.google.com/go/storage
RUN go get google.golang.org/api/
COPY ./downloader .
RUN go build -o downloader .



# build reloader
FROM golang:stretch as reloader
ENV GOPATH /build
WORKDIR /build
RUN go get github.com/NeoJRotary/exec-go
COPY ./reloader .
RUN go build -o reloader .



# build main nginx
FROM nginx:1.16.0
RUN apt-get update -y
RUN apt-get install bash curl -y
RUN rm /etc/nginx/conf.d/default.conf
# default ENV
ENV ENABLE_GCS_SYNC false
ENV GOOGLE_APPLICATION_CREDENTIALS /service_account.json
ENV CONFIGMAP_FOLDER /etc/config
ENV SHOW_RELOAD_CHECK_RESULT false

COPY --from=downloader /build/downloader /downloader
COPY --from=reloader /build/reloader /reloader
COPY shells/. /
COPY module-conf /module-conf
CMD /start.sh
