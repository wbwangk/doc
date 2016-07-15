---
layout: page
title: 平台概述
---
# 平台概述 #

## 租户注册实现过程
1. 用户进入平台首页，点击注册按钮 
2. 用户输入邮箱、用户名、密码等，IAM创建用户
3. 租户登陆，回到平台首页
4. 首页上集成了管理控制台ADM的小窗口，小窗口显示用户信息：二级域名:（无）、github账号:（无）。小窗口上有进入ADM的按钮
5. 用户点击按钮进入管理控制台，控制台上仍显示“二级域名:（无）、github账号:（无）”，但两个“（无）”都成了超链接，点击超链接可以触发出输入框供用户输入。
6. 二级域名即租户的用户id，用ajax把域名信息写入etcd（在etcd中创建```/nginx/vhosts/$subdomain/```目录。然后在目录下创建一个"github-id"的key，值初始化为"imaidev"），输入框消失。github账号同理，用户可以选择不输入github账号。
7.confd监听着etcd的数据变化，自动生成nginx的配置文件。如果没有租户定义自己的github账号，则github-id的值是imaidev,所以静态内容反向代理到imaidev.github.io。如果租户定义了自己的github账号，则静态内容反向代理到"$github-id.github.io"。

## API Key申请和使用的实现过程
1. 用户通过ADM进入API key管理界面，显示有效API key清单
2. 点击创建按钮，系统生成一对随机数（分别是api key id和api key secret）作为api key保存到etcd。使用apache htpaaswd算法生成api key密码文件，密码文档保存到etcd中租户名下。
3. 系统生成一个api key文件，并自动下载。提示用户保存好api key文件
4. confd监听etcd中api key的数据变化，并将各个租户的htpasswd密码文件放置到nginx的指定目录下。conf还要同步etcd中其他数据
5. nginx重新加载配置（reload）
6. 用户（开发者）使用curl（或写程序）调用平台API，curl写法见备注
7. nginx的API服务相关location中已经配置http基础认证，当nginx检测到api服务请求时会检查http头中用户名密码与htpasswd文件中是否一致

备注：基于NginX的API认证(Basic Authentication)。通过API Key来认证API服务。API Key放置在Http Header中。例如：
```
curl --request GET \
--user $API_KEY_ID:$API_KEY_SECRET \
--header 'content-type: application/json' \
--url "https://api.stormpath.com/v1/tenants/current"
```

## etcd中数据结构

- ```/subdomain-list/```   存放备选的二级域名
- ```/nginx/```  存放了与nginx配置文件有关的元数据
- ```/nginx/vhosts/``` 存放了所有的虚拟主机
- ```/nginx/vhosts/$subdomain/``` 使用二级域名当vhost id，即$subdomain.imaicloud.com，如a001.imaicloud.com
- ```/nginx/vhosts/$subdomain/github-id``` 租户的github账号，默认值是imaidev
- ```/nginx/vhosts/$subdomain/apikeys/``` 存放有效的api key，key是htpasswd中的user id。
