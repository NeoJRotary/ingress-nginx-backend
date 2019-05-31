## v1.2.0
- Add `SHOW_RELOAD_CHECK_RESULT` ENV to enable showing ConfigMap reload checking's result
- Rebuild image for latest version (1.16.0) and move to debian/stretch base

## v1.1.1
- Fix Reloader's output issue ( print full content for debugging )   

## v1.1.0
- Support Kubernetes ConfigMap
- Support Nginx hot reload from ConfigMap
- Function `GCS Sync` becomes selectable
- Update ENV names