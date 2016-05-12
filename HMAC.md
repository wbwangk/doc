---
layout: page
title: Hash Message Authentication Code(HMAC)
permalink: /hmac/
---

算法用到的java类是javax.crypto.Mac  
A MAC provides a way to check the integrity of information transmitted over or stored in an unreliable medium, based on a secret key. Typically, message authentication codes are used between two parties that share a secret key in order to validate information transmitted between these parties.  

A MAC mechanism that is based on cryptographic hash functions is referred to as HMAC. HMAC can be used with any cryptographic hash function, e.g., MD5 or SHA-1, in combination with a secret shared key. HMAC is specified in RFC 2104.  

Every implementation of the Java platform is required to support the following standard Mac algorithms:

- HmacMD5
- HmacSHA1
- HmacSHA256
