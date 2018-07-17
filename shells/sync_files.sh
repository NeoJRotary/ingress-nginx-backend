#!/bin/bash
set -e

rm -rf /etc/nginx/conf.d/*

if [ "$ENABLE_GCS_SYNC" = "true" ]
then
  /downloader
fi

cp -R $CONFIGMAP_FOLDER/. /etc/nginx/conf.d/

if [ -f "/etc/nginx/conf.d/nginx.conf" ]
then
  rm /etc/nginx/nginx.conf
  mv /etc/nginx/conf.d/nginx.conf /etc/nginx/nginx.conf
fi

/enable_modules.sh

if [ -f "/etc/nginx/conf.d/before.sh" ]
then
  mv /etc/nginx/conf.d/before.sh /before.sh
  /before.sh
fi