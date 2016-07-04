---
layout: page
title: 平台概述
---
# 平台概述 #

## 租户注册实现过程
1. 用户进入平台首页，点击注册按钮 
2. 用户输入邮箱、用户名、密码等，IAM创建用户
3. 租户登陆，回到平台首页
4. 首页上集成了管理控制台ADM的小窗口，小窗口显示用户信息：```二级域名:（无）、启用服务:0个```。小窗口上有进入ADM的按钮
5. 用户点击按钮进入管理控制台，控制台上仍显示```二级域名:（无）、启用服务:0个```，但两个都成了超链接。紧挨着下方显示```已申请的API Key: 0个```。
6. 用户输入二级域名（当前实现是从etcd的```subdomain-list```目录中选），如果域名已经被占用就提示用户。用ajax把域名信息写入etcd（删除```subdomain-list```中选定的记录，并在etcd中创建```/nginx/vhosts/$subdomain/```目录。
7. 用户点击“已启用的组件服务”进入组件服务管理界面。（全部组件列表存放在etcd的```front-list```目录下。）用户点击“新增”按钮，从尚未启用的组件中选择一个。（组件服务启用后在```/nginx/vhosts/$subdomain/fronts```目录下会增加记录。）用户可以将组件的前端反向代理到自己的github.io个性地址（如```gh-a001.github.io/ocs```）。
8. confd监听着etcd的数据变化，自动生成nginx的配置文件。nginx配置文件的内容有server_name,值是etcd目录```$subdomain```。nginx配置中还会生成多个location，启用了几个服务就生成几个location。location的proxy_pass值就是```fronts```目录下各个key的value。

## API Key申请和使用的实现过程
1. 用户通过ADM进入API key管理界面，显示有效API key清单
2. 点击创建按钮，系统生成一对随机数（分别是api key id和api key secret）作为api key保存到etcd（初期用脚本触发htpasswd工具生成？）
3. 系统生成一个api key文件，并自动下载。提示用户保存好api key文件
4. confd监听etcd中api key的数据变化，并自动生成htpasswd文件。
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

- ```/apikey-list/```   存放备选的api key(使用[htpasswd](https://en.wikipedia.org/wiki/.htpasswd))，key是htpasswd中的user id。
- ```/subdomain-list/```   存放备选的二级域名
- ```/nginx/```  存放了与nginx配置文件有关的元数据
- ```/nginx/vhosts/``` 存放了所有的虚拟主机
- ```/nginx/vhosts/$subdomain/``` 使用二级域名当vhost id，即$subdomain.imaicloud.com，如a001.imaicloud.com
- ```/nginx/vhosts/$subdomain/github-id``` 租户的github账号，默认值是imaidev
- ```/nginx/vhosts/$subdomain/apikeys/``` 存放有效的api key，key是htpasswd中的user id。

