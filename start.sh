#!/bin/bash
set -e

echo "$SERVICE_ACCOUNT" > "/service_account.json"

# download files
/initConfig

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

nginx -g "daemon off;"