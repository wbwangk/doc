---
layout: pages
title: 免费https证书(Let's Encrypt)
---

letsencrypt.org可以发放免费的https证书。

## 原理

需要在webserver的机器上安装一个代理，推荐[cetbot](https://certbot.eff.org/)。cetbot执行后，会在webserver的webroot目录下创建隐藏的目录和文件。如nginx如果安装在/opt目录下，root指定为html，则cetbot会在/opt/nginx/html目录下创建/.well-known/acme-challenge的网页。let's encrypt的官网会调用这个网页 ，目的是查看certbot所在的webserver是否对于要发放证书的域名有控制权。
