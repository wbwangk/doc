---
layout: page
title: 域名与路径
permalink: /domain/
---
## 二级域名 ##
imaicloud.com的二级主要供租户使用，如wbwang.imaicloud.com标识租户wbwang在本站托管的资源，主要是静态文件。部分二级域名是保留的，如dev.imaicloud.com是imai云的开发者中心。

## 代理服务器上的目录规划 ##



组件SaaS暂不使用二级域名，而使用目录，如imaicloud.com/ocs是在线客服组件的根路径，也是静态文件存放的目录。各个war包的上下文根需要和其静态文件所在目录类似，如在线客服的war包可以用类似ocs_1或ocs_web当上下文根。

假定各组件的静态文件是开放的，各个租户可以定制


