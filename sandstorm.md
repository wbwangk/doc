---
layout: page
title: sandstorm研究
permalink: /sandstorm/
---

sandstorm平台负责用户认证,应用不用管。而且sandstorm要求应用不能实现用户认证。

# 一般用户请求
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

# API请求
sandstorm会检查HTTP请求头中的API token，如果有效就移除头中token，然后在http头中增加一些属性，如X-Sandstorm-Username。当应用的http响应返回后，sandstorm会在响应头中再增加一些内容，然后再发给请求方。
例子：
```
curl -H "Authorization: Bearer 49Np9sqkYV4g_FpOQk1p0j1yJlvoHrZm9SVhQt7H2-9" https://alpha-api.sandstorm.io/
```
上面的格式是OAuth 2.0-style Bearer header

### websocket请求
websocket请求无法设置http头，只能把API放在URL中，格式是：
```
/.sandstorm-api-token/<token>
```
例子：
```
wss://api-qxJ58hKANkbmJLQdSDk4.oasis.sandstorm.io/.sandstorm-api-token/RfNqni4FEHXkWC5B8v6t/some/path
```
 "/.sandstorm-api-token/<token>" 部分会被sandstorm剥离，发往应用的是剩余部分。

### AJAX跨域请求
Sandstorm applies a CORS header of ```Access-Control-Allow-Origin: * ``to allow Javascript on any domain to interact with the app's API. This is safe because the API token serves as the access control.
