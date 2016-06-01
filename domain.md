---
layout: page
title: 域名与路径
permalink: /domain/
---
## 二级域名 ##
`imaicloud.com`的二级主要供租户使用，如`lucy.imaicloud.com`标识租户`lucy`在本站托管的资源，主要是静态文件。有些租户是网站官方的，如`dev`是imai云的开发者中心租户，开发者中心的网址即`dev.imaicloud.com`。

## 代理服务器上的目录规划 ##
每个二级域名（即每个租户）在代理服务上对应一个`vhost`，如`dev.imaicloud.com`或`wbwang.imaicloud.com`都是`vhost`。`vhost`下每个组件都有对应的目录存放各自的静态文件，称组件目录。如ocs是在线客服组件的组件目录，`dev.imaicloud.com/ocs/xxx.js`是在线客服组件的某文件url。同一个文件，在不同租户下可能有个性化版本，如`lucy.imaicloud.com/ocs/xxx.js`是租户`lucy`下的某js文件的个性化版本。

## 各组件war的命名规划 ##
每个组件可能有多个war。各个war包的命名(上下文根)需要和其静态文件所在目录一致。如在线客服的war包可以用类似ocs_1或ocs_web来命名。war不开源，也不提供个性化定制。

## github上的库(repository)规划
在github上建一个组织对应开发者中心(`dev.imaicloud.com`)，目前组织叫`imaip`。在`imaip`组织下，每个组件都有自己的库。如`ocs`的库叫`ocs`。


