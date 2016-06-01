---
layout: page
title: 域名与路径
permalink: /domain/
---
## 二级域名 ##
`imaicloud.com`的二级主要供租户使用，如`wbwang.imaicloud.com`标识租户`wbwang`在本站托管的资源，主要是静态文件。有些租户是网站官方的，如`dev`是imai云的开发者中心租户，开发者中心的网址即`dev.imaicloud.com`。

## 代理服务器上的目录规划 ##

每个二级域名（即每个租户）在代理服务上对应一个`vhost`，如`dev.imaicloud.com`或`wbwang.imaicloud.com`都是`vhost`。每个组件

组件SaaS暂不使用二级域名，而使用目录，如imaicloud.com/ocs是在线客服组件的根路径，也是静态文件存放的目录。各个war包的上下文根需要和其静态文件所在目录类似，如在线客服的war包可以用类似ocs_1或ocs_web当上下文根。

假定各组件的静态文件是开放的，各个租户可以定制


