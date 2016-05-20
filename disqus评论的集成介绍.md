---
layout: page
title: disqus研究
permalink: /disqus/
---

[本文的来源](https://help.disqus.com/customer/en/portal/topics/107054-developer/articles)

##注册disqus##
首先，需要你需要在disqus注册一个用户。你将使用这个用户登录disqus来管理你的网站。在disqus中注册你的网站时，需要为你的网站选择一个shortname来标识你的网站。disqus会用这个shortname自动注册二级域名。  

##通用嵌入代码##

首先，需要到disqus注册你的网站。 通过 [Quickstart Guide](https://help.disqus.com/customer/portal/articles/466182-quick-start-guide) 过去更多信息。
然后，你需要知道注册时填写的 forum shortname。
disqus嵌入javasrcript代码如下：  

```
<div id="disqus_thread"></div>
<script>
    /**
     *  RECOMMENDED CONFIGURATION VARIABLES: EDIT AND UNCOMMENT THE SECTION BELOW TO INSERT DYNAMIC VALUES FROM YOUR PLATFORM OR CMS.
     *  LEARN WHY DEFINING THESE VARIABLES IS IMPORTANT: https://disqus.com/admin/universalcode/#configuration-variables
     */
    /*
    var disqus_config = function () {
        this.page.url = PAGE_URL;  // Replace PAGE_URL with your page's canonical URL variable
        this.page.identifier = PAGE_IDENTIFIER; // Replace PAGE_IDENTIFIER with your page's unique identifier variable
    };
    */
    (function() {  // REQUIRED CONFIGURATION VARIABLE: EDIT THE SHORTNAME BELOW
        var d = document, s = d.createElement('script');
        
        s.src = '//EXAMPLE.disqus.com/embed.js';  // IMPORTANT: Replace EXAMPLE with your forum shortname!
        
        s.setAttribute('data-timestamp', +new Date());
        (d.head || d.body).appendChild(s);
    })();
</script>
<noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript" rel="nofollow">comments powered by Disqus.</a></noscript>
```  

**EXAMPLE**是注册到disqus的shortname，告诉disqus这些评论属于哪个“租户”。
**this.page.url**告诉disqus当前页面的url。
**this.page.identifier**告诉disqus当前页面的id，这样如果URL改变id可以不变。这是一个可选变量，但建议用。如果不舍定这个变量，则会默认使用page.url当identifier。



## disqus单点登录 ##
  使用外部系统的用户名口令登录disqus而不用注册disqus用户，前提是你已经有了一个用户库。*如果你有多个网站，每个网站都要注册一个远程域。*
  
  你的网站需要增加一个disqus插件，资格也需要手工向disqus申请（发邮件？）。
  
  插件工作原理：用户登录你的网站后，会根据用户凭据动态生成一个payload发给disqus。这样允许用户仅登录你的网站就能参与disqus讨论。
  
  disqus SSO会基于你的网站为你用户建立档案（profile）。猜测是网站id+用户id之类的方法。这样做为了防止SSO用户与disqus已有用户的id冲突。
  
  首先，需要你进入SSO页面去配置远程域。
  
  要生成一个令牌传递给disqus。令牌分三部分：消息、签名、时间戳，即：
  ```
  this.page.remote_auth_s3 = '<message> <hmac> <timestamp>';
  ```
  消息的构成：id、username、email以及两个可选参数avatar（用户头像url）、url（用户个人网站url）。id要保证在disqus内的唯一性。
  签名算法：``` HMAC->SHA1(secret_key, message + ' ' + timestamp) ```
  
### 增加自己的登录登出链接 ###
例子：
```
var disqus_config = function () {
    // The generated payload which authenticates users with Disqus
    this.page.remote_auth_s3 = '<message> <hmac> <timestamp>';
    this.page.api_key = 'public_api_key';

// This adds the custom login/logout functionality
    this.sso = {
          name:   "SampleNews",
          button:  "http://example.com/images/samplenews.gif",
          icon:     "http://example.com/favicon.png",
          url:        "http://example.com/login/",
          logout:  "http://example.com/logout/",
          width:   "800",
          height:  "400"
    };
};
```

**说明：**

- **name** — 你的网站名称。 We will display it in the Post As window.
- **button** — 按钮图片的地址。 Disqus 2012 users, see style guide below.
- **icon** — Address of the image that appears on the login modal's SSO tab. Favicons work well here. (Not required in Disqus 2012.)
- **url** — 你的登录页地址。这个页面将打开新窗口，这个页面必须在认证后自己关闭。 That's how we know when it is done and reload the page.
- **logout** — 你的登出页面。你的用户点登出按钮后被导向到页面。
- **width** — 登录窗口宽度，默认800。
- **height** — 登录窗口高度，默认400。
以上的所有URL都必须是绝对链接，因为它们将在disqus的iframe中加载。

**创建自己的登录按钮**
1.下载[SSO login button style guide](http://content.disqus.com/design/disqus-sso-login-button-template.psd)
2.用你自己的logo覆盖"Replace this logo"层
3.修改按钮颜色("Button"层)为符合你的品牌
4.保持形状的边框在预设范围内，以便与其他图标匹配

如果你的登录窗口不能自动关闭，请使用javascript脚本自动关闭它。


{% for i in (1..5) %}
i= {{ i }} 
{% endfor %}

