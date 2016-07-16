---
layout: page
title: 免费https证书(Let's Encrypt)
---

letsencrypt.org可以发放免费的https证书。

## 原理

需要在webserver的机器上安装一个代理，推荐[cetbot](https://certbot.eff.org/)。
```
$ sudo yum install epel-release
$ sudo yum install certbot
```
执行cetbot：```$ sudo certbot certonly```。会弹出一个导航窗口。选择webroot。
1. 
会在webserver的webroot目录下创建隐藏的目录和文件。如nginx如果安装在/opt目录下，root指定为html，则cetbot会在```/opt/nginx/html```目录下创建```/.well-known/acme-challenge```的网页。let's encrypt的官网会调用这个网页 ，目的是查看certbot所在的webserver是否对于要发放证书的域名有控制权。

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

## 附录 ##
为多个子域名一起申请证书:
certbot certonly --webroot -w /opt/nginx/html-imaicloud/ -d imaicloud.com -d www.imaicloud.com -d dockerui.imaicloud.com -d dev.imaicloud.com -d ethercalc.imaicloud.com -d etcd.imaicloud.com -d couchdb.imaicloud.com -d uba.imaicloud.com -d registry.imaicloud.com

### OpenSSL
Mozialla的trust store：https://hg.mozilla.org/mozilla-central/raw-file/tip/security/nss/lib/ckfw/builtins/certdata.txt（专属格式）
CURL有一个PEM格式的镜像库：http://curl.haxx.se/docs/caextract.html

Most users turn to OpenSSL because they wish to configure and run a web server that supports SSL. That process consists of three steps: (1) generate a strong private key, (2) create a Certificate Signing Request (CSR) and send it to a CA, and (3) install the CA-provided certificate in your web server.

Creating a new CA involves several steps: configuration, creation of a directory structure and initialization of the key files, and finally generation of the root key and certificate. This section describes the process as well as the common CA operations.
