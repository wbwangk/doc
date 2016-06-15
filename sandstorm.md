---
layout: page
title: sandstorm研究
permalink: /sandstorm/
---

sandstorm平台负责用户认证,应用不用管。而且sandstorm要求应用不能实现用户认证。

sandstorm会过滤所有的http请求。http请求会通过一个叫`sandstorm-http-bridge`的工具转给应用:
 - 在http头上增加标记来标识当前用户("authentication")
 - 在http头上增加权限标记("authorization")
 - 将一些不合规的请求拦截掉
 
 经过sandstorm-http-bridge的处理，应用收到的http头有：
 - X-Sandstorm-Username：用户名
 - X-Sandstorm-User-Id：用户id，是用户id的SHA256编码的头128个bit，如0ba26e59c64ec75dedbc11679f267a40
 - X-Sandstorm-Tab-Id：当前操作的grain选项卡id
 - X-Sandstorm-Permissions：权限清单
 - X-Sandstorm-Preferred-Handle：参数或偏好？
 - X-Sandstorm-User-Picture：用户头像url。也可以用https://github.com/stewartlord/identicon.js来为用户生成一个(X-Sandstorm-User-Id当输入)
 - X-Sandstorm-User-Pronouns：they/he/him/her/she/it等，代称

也可以不使用`sandstorm-http-bridge`，然后使用服务访问认证信息。
