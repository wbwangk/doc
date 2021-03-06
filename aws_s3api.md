---
layout: page
title: 对请求的认证(AWS Signature Version 4)
permalink: /s3api/
---

同S3的交互要么是认证的，要么是匿名的。本章解释如何使用亚马逊V4版本签名算法对请求进行认证。如果使用SDK则不用阅读本章，因为SDK客户端使用你提供的access key来认证请求。
亚马逊V4签名提供的认证有：

- **验证请求者的身份** 认证请求需要一个签名，签名用你的access key（access key id和secret access key）生成。

- **传输数据保护** 为了防止传输数据被篡改，使用部分关键传输数据进行签名

- **方式重放攻击** 被签名的内容中有一个15分钟过期的时间戳

### 认证方法 ###

- **http授权头** 在http请求头中放置Authorization是常用的认证S3请求的方法。除了基于浏览器的上传POST，所有S3 REST操作都需要这个头。
- **请求参数** 签名以参数的形式放置在URL中。
- **基于浏览器的POST上传(AWS Signature Version 2)** 

下图解释了AWS S3 POST：

![](http://docs.aws.amazon.com/AmazonS3/latest/dev/images/s3_post.png)

使用表单来上传文件,表单头:
```
    <form action="http://johnsmith.s3.amazonaws.com/" method="post"
    enctype="multipart/form-data">
``` 

### 签名请求介绍 ###
你发送的认证请求中包含一个签名。AWS不是直接使用secret access key来签名，而是使用secret access key先生成一个sign key。

![](http://docs.aws.amazon.com/AmazonS3/latest/API/images/signing-overview.png)

（上图表达的sign key生成逻辑与stormpath的摘要认证算法类似，特征是多次散列迭代）

AWS S3的签名算法很复杂，在应用程序中直接写很麻烦，可以靠下面两种办法代替手写程序：
- 使用AWS SDK，里面内置了签名算法
- 使用AWS [CLI](http://docs.aws.amazon.com/cli/latest/userguide/cli-s3.html) 发起API调用
