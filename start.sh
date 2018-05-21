#!/bin/bash
set -e

echo "$SERVICE_ACCOUNT" > "/service_account.json"
/initConfig
if [ -f "/etc/nginx/conf.d/nginx.conf" ]
then
  rm /etc/nginx/nginx.conf
  mv /etc/nginx/conf.d/nginx.conf /etc/nginx/nginx.conf
fi
nginx -g "daemon off;"