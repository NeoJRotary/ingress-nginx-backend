#!/bin/bash
set -e

NGINXCONF="/etc/nginx/nginx.conf"
CONFD="/etc/nginx/conf.d"

if [ "$ENABLE_MODULE_GEOIP" = "true" ]
then

# download 
geoipDir="/usr/share/GeoIP"
mkdir -p $geoipDir
curl http://geolite.maxmind.com/download/geoip/database/GeoLiteCountry/GeoIP.dat.gz --output $geoipDir/GeoIP.dat.gz
gunzip $geoipDir/GeoIP.dat.gz
curl http://geolite.maxmind.com/download/geoip/database/GeoLiteCity.dat.gz --output $geoipDir/GeoLiteCity.dat.gz
gunzip $geoipDir/GeoLiteCity.dat.gz

# add load_module
config="load_module \"modules/ngx_http_geoip_module.so\";"
echo "$(echo $config | cat - $NGINXCONF)" > $NGINXCONF

# cp config
cp /module-conf/geoip.conf $CONFD/geoip.conf

fi
