# imai云开发平台


### 基于NginX的API认证(Basic Authentication)
通过API Key来认证API服务。API Key放置在Http Header中。例子：
```
curl --request GET \
--user $SP_API_KEY_ID:$SP_API_KEY_SECRET \
--header 'content-type: application/json' \
--url "https://api.stormpath.com/v1/tenants/current"
```

## 自举过程
1. NginX验证API Key调试通过
3. 通过NginX+LUA实现了Vhost的创建REST服务。
4. IAM调试通过。
5. ADM结合IAM租户注册和Vhost的创建服务,实现平台租户注册的完整过程。（可以API Key或防火墙保护，避免vhost创建服务滥用）
6. 租户通过IAM创建API Key，Key被复制到NginX上，放置到租户目录根上。
7. 
