#!/bin/bash
set -e

echo "$SERVICE_ACCOUNT" > "/service_account.json"

/reloader &

# download files
/sync_files.sh

nginx -g "daemon off;"