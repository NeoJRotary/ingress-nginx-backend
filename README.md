# Ingress NGINX Backend v1.1.0
Kubernetes Ingress NGINX Backend

## Intro
- Test with kubernetes 1.9.6   
- Check `CHANGESmd` for updates   
- You can get it from Docker Hub   
`docker pull neojrotary/ingress-nginx-backend`

## How To Use
Setup ENVs, Done!    
It will copy all files in your Cloud Storage / ConfigMap Volume into `/etc/nginx/conf.d/`.

### Special Filename
Below filenames in the folder have different behavior.
- `nginx.conf`   
It will overwrite `/etc/nginx/nginx.conf`
- `before.sh`   
It will be execute at root folder to help you setup something before start nginx server. For example, curl to download geoip files for geoip nginx module.

## ENV
### Cloud Storage Provider
- ENABLE_GCS_SYNC : enable GCS sync function. Default is false.
- SERVICE_ACCOUNT : your service account credential content in json format
- BUCKET_NAME : bucket name
- BUCKET_FOLDER : config files folder name in the bucket. It will download all files in it.

### Nginx Module
- ENABLE_MODULE_GEOIP : if `true` it will download `Maxmind GeoLite` and setup GeoIP module for you.

### K8S ConfigMap
- CONFIGMAP_FOLDER : folder which mount with kubernetes configmap. Default is `/etc/config/`.
- CONFIGMAP_SCAN_DUR : duration between scanning of configmap in second. Default is 60s.

## Setup Service Account
You can set value directly at k8s configuration or pass by k8s secret.

Create secret : 
```
apiVersion: v1
kind: Secret
metadata:
  name: service-account
  namespace: default
type: Opaque
data:
  storage-only: [my service account]
```

In your container spec :
```
env:
  - name: SERVICE_ACCOUNT
    valueFrom:
      secretKeyRef:
        name: service-account
        key: storage-only
```
