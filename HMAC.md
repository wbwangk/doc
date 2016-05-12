---
layout: page
title: Hash Message Authentication Code(HMAC)
permalink: /hmac/
---

算法用到的java类是[javax.crypto.Mac](http://docs.oracle.com/javase/7/docs/api/javax/crypto/Mac.html)  

A MAC provides a way to check the integrity of information transmitted over or stored in an unreliable medium, based on a secret key. Typically, message authentication codes are used between two parties that share a secret key in order to validate information transmitted between these parties.  

A MAC mechanism that is based on cryptographic hash functions is referred to as HMAC. HMAC can be used with any cryptographic hash function, e.g., MD5 or SHA-1, in combination with a secret shared key. HMAC is specified in RFC 2104.  

Every implementation of the Java platform is required to support the following standard Mac algorithms:

- HmacMD5
- HmacSHA1
- HmacSHA256

在stormpath中使用Mac这个类进行摘要认证，使用源码可以参照stormpath SDK类[SAuthc1RequestAuthenticator.java][1]  
通过分析类[JwtWrapper.java][2]可知，JWT由dot隔开的三部分构成：``` <jwt头>.<jwt负载(payload)>.<签名> ```  
通过分析类[HmacGenerator.java][3]可以看到jwt的生成方式，代码不多。  


[1]: https://github.com/stormpath/stormpath-sdk-java/blob/19dbc0a9b811c427a8863609658947cffd6fbd26/impl/src/main/java/com/stormpath/sdk/impl/http/authc/SAuthc1RequestAuthenticator.java
[2]: https://github.com/stormpath/stormpath-sdk-java/blob/master/impl/src/main/java/com/stormpath/sdk/impl/jwt/JwtWrapper.java
[3]: https://github.com/stormpath/stormpath-sdk-java/blob/master/impl/src/main/java/com/stormpath/sdk/impl/jwt/signer/HmacGenerator.java
