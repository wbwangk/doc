---
layout: pages
title: 免费https证书(Let's Encrypt)
---

letsencrypt.org可以发放免费的https证书。

## 原理

需要在webserver的机器上安装一个代理，推荐[cetbot](https://certbot.eff.org/)。cetbot执行后，会在webserver的webroot目录下创建隐藏的目录和文件。如nginx如果安装在/opt目录下，root指定为html，则cetbot会在```/opt/nginx/html```目录下创建```/.well-known/acme-challenge```的网页。let's encrypt的官网会调用这个网页 ，目的是查看certbot所在的webserver是否对于要发放证书的域名有控制权。

ertbot还会在```/etc/letsencrypt/live/```目录下创建目录存放数字证书。如：
```
/etc/letsencrypt/live/imaicloud.com/fullchain.pem
/etc/letsencrypt/live/imaicloud.com/privkey.pem
```
然后配置nginx的配置文件：
```
server {
    listen       443;
    server_name  imaicloud.com;
    root   html;
    ssl                  on;
    ssl_certificate      /etc/letsencrypt/live/imaicloud.com/fullchain.pem;
    ssl_certificate_key  /etc/letsencrypt/live/imaicloud.com/privkey.pem;
```
nginx重新装载配置文件。然后就可以用浏览器访问```https://imaicloud.com```来测试了。

## 过程（CentOS为例）
安装certbot：
```
$ sudo yum install epel-release
$ sudo yum install certbot
```
执行certbot：
```
$ sudo certbot certonly
```
会弹出一个简陋的导航窗口。选择```webroot```，之后输入nginx的web root目录（如``/opt/nginx/html```）。屏幕会提示英文的“恭喜”，以及产生的证书。
修改nginx配置文件，并重新装载。
