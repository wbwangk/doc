---
layout: page
title: disqus研究
permalink: /disqus/
---

## disqus单点登录 ##
  使用外部系统的用户名口令登录disqus而不用注册disqus用户，前提是你已经有了一个用户库。如果你有多个网站，每个网站都要注册一个远程域。
  
  你的网站需要增加一个disqus插件，资格也需要手工向disqus申请（发邮件？）。
  
  插件工作原理：用户登录你的网站后，会根据用户凭据动态生成一个payload发给disqus。这样允许用户仅登录你的网站就能参与disqus讨论。
  
  disqus SSO会基于你的网站为你用户建立档案（profile）。猜测是网站id+用户id之类的方法。这样做为了防止SSO用户与disqus已有用户的id冲突。
  
  首先，需要你进入SSO页面去配置远程域。
  
  要生成一个令牌传递给disqus。令牌分三部分：消息、签名、时间戳，即：
  ···
  this.page.remote_auth_s3 = '<message> <hmac> <timestamp>';
  ···
  消息的构成：id、username、email以及两个可选参数avatar（用户头像url）、url（用户个人网站url）。id要保证在disqus内的唯一行。
  签名算法：··· HMAC->SHA1(secret_key, message + ' ' + timestamp) ··· 
