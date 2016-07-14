---
layout: page
title: 平台容器的启动脚本
---

## 启动couchdb(couch/fo) 7.105
```
docker run -p 5984:5984 -d -v /var/couchdb:/usr/local/var/lib/couchdb \
 -v /var/couchdb/conf:/usr/local/etc/couchdb/local.d  \
 -v /var/couchdb/log:/usr/local/var/log/couchdb \
registry.aliyuncs.com/imaidev/couchdb
```

## 启动ui-for-docker（docker容器管理）
```
docker run -d -p 9001:9000 --privileged -v /var/run/docker.sock:/var/run/docker.sock registry.aliyuncs.com/imaidev/ui-for-docker
```
## 启动docker镜像库（docker registry）
```
docker run -d -p 5000:5000 --restart=always -v /var/lib/registry:/var/lib/registry \
--name registry  registry.aliyuncs.com/imaidev/registry
```
## 启动etcd-viewer
```
docker run -d -p 8080:8080 registry.aliyuncs.com/imaidev/etcd-viewer
```
## 启动tobegit3hub/seagull（容器管理，界面主题多）
```
docker run -d -p 10086:10086 -v /var/run/docker.sock:/var/run/docker.sock registry.imaicloud.com/tobegit3hub/seagull
```


