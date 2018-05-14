# Ingress NGINX Backend
Kubernetes Ingress NGINX Backend with GCP

## Docker Hub
`docker pull neojrotary/ingress-nginx-backend`

## How To Use
Setup ENV, Done!   
Everytime container start it will download latest config from your GCP bucket then start nginx.

## ENV
- SERVICE_ACCOUNT : your service account credential content in json format
- GCS_BUCKET : bucket name
- CONFIG_FOLDER : config files folder. It will download all files in it.

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
