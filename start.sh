#!/bin/bash
set -e

echo "$SERVICE_ACCOUNT" > "/service_account.json"
/initConfig
nginx -g "daemon off;"