# imai云开发平台


### 基于NginX的API认证(Basic Authentication)
通过API Key来认证API服务。API Key放置在Http Header中。例子：
```
curl --request GET \
--user $SP_API_KEY_ID:$SP_API_KEY_SECRET \
--header 'content-type: application/json' \
--url "https://api.stormpath.com/v1/tenants/current"
```

## 租户注册实现过程
1. 用户进入平台首页，点击注册按钮 
2. 用户输入邮箱、用户名、密码等，IAM创建用户和租户
3. 租户登陆，回到平台首页
4. 首页上集成了管理控制台ADM的小窗口，小窗口提示用户信息：虚拟主机（无）、github账号（无）。小窗口上有进入ADM的按钮
5. 用户点击按钮进入管理控制台，控制台上仍显示“虚拟主机（无）、github账号（无）”，但两个“（无）”都成了超链接，点击超链接可以触发出输入框供用户定义。
6. 用户定义虚拟主机（系统自动帮他选一个二级域名），用ajax把域名、虚拟主机信息写入etcd，输入框消失。github账号同理，用户可以选择不输入github账号。
7.confd监听着etcd的数据变化，自动生成nginx的配置文件。对于vhost，如果没有定义github账号，则静态内容反向代理到dev.imaicloud.com。如果定义了github账号，则静态内容反向代理到个性github库(库名按约定)。

## API Key的定义过程
