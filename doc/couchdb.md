---
layout: page
title: couchdb的安全模块
---

## 基础认证（--user可以简写为-u）
1. ```curl --user anna:secret  https://couchdb.imaicloud.com/_config```
2. ```curl https://anna:secret@couchdb.imaicloud.com/_config```

以上两种url用法证明都可以。第2种写法在浏览器地址栏中输入是不行的。

创建一个叫somedatabase的数据库：```curl -X PUT --user anna:secret https://couchdb.imaicloud.com/somedatabase```

## cookie认证
登陆：(-v参数会显示http调用过程的详细信息)
```
 curl -vX POST http://10.0.7.105:5984/_session \
       -H 'Content-Type:application/x-www-form-urlencoded' \
       -d 'name=anna&password=secret'
```
或
```
 curl -vX POST https://dev.imaicloud.com/signin \
       -H 'Content-Type:application/x-www-form-urlencoded' \
       -d 'name=anna&password=secret'
```

增加一个数据库，测试cookie认证
```
 curl -vX PUT http://10.0.7.105:5984/mydatabase \
       --cookie AuthSession=YW5uYTo1Nzg2MzU4RTqigPm9-aUOs2Q7qaBZcTRvvOUCHg \
       -H "X-CouchDB-WWW-Authenticate: Cookie" \
       -H "Content-Type:application/x-www-form-urlencoded"
```
创建用户：
```
curl -X PUT http://10.0.7.105:5984/_users/org.couchdb.user:jan \
     -H "Accept: application/json" \
     -H "Content-Type: application/json" \
     -d '{"name": "jan", "password": "apple", "roles": [], "type": "user"}'
```
登陆：
```
curl -vX POST http://10.0.7.105:5984/_session -d 'name=jan&password=apple'
```
响应(Set-Cookie就是couchdb创建的cookie)：
```
< HTTP/1.1 200 OK
< Set-Cookie: AuthSession=amFuOjU3ODY0NzU1Otyly0ka4T1Y5FGB0Q8yfZGfmvbq; Version=1; Path=/; HttpOnly

```
取用户信息（--cookie参数模拟了请求中cookie的值）：
```
curl -vX GET http://10.0.7.105:5984/_users/org.couchdb.user:jan \
      --cookie AuthSession=amFuOjU3ODY0NzU1Otyly0ka4T1Y5FGB0Q8yfZGfmvbq
```
响应：
```{"_id":"org.couchdb.user:jan","_rev":"1-73f16884aa08df40b1a1b1710d1dcad2","password_scheme":"pbkdf2","iterations":10,"name":"jan","roles":[],"type":"user","derived_key":"4721471dfe82ee266d493acf26d5f8927d916435","salt":"2011ade37425b921ef87aec9f0c1daa5"}
```
改密码：
```
curl -X PUT http://10.0.7.105:5984/_users/org.couchdb.user:jan \
     -H "Accept: application/json" \
     -H "Content-Type: application/json" \
     -H "If-Match: 1-73f16884aa08df40b1a1b1710d1dcad2" \
      --cookie AuthSession=amFuOjU3ODY0NzU1Otyly0ka4T1Y5FGB0Q8yfZGfmvbq \
     -d '{"name":"jan", "roles":[], "type":"user", "password":"orange"}'     （If-Match填入上面响应的_rev）
```
-- 注：couchdb的官方文档中都没有加--cookie的参数，估计curl可以自动保存cookie，但我的不行，只能手工加上这个--cookie。

试一下新密码：```curl -X POST http://10.0.7.105:5984/_session -d 'name=jan&password=orange'```
响应：```{"ok":true,"name":"jan","roles":[]}```

```
curl -X PUT http://10.0.7.105:5984/_config/couch_httpd_auth/public_fields \
-H “Content-Type: application/json” -d ‘“name”’ -u admin   （提示输入密码，但密码忘记了）
```

之前创建用户的请求：```curl -X PUT 10.0.7.105/_config/admins/anna -d '"secret"'```创建的用户是服务器管理员（server admin）。在没有数据库管理员前，她也被视为数据库管理员。
```
curl -X PUT http://10.0.7.105:5984/mydatabase/_security \
     -u anna:secret \
     -H "Content-Type: application/json" \
     -d '{"admins": { "names": [], "roles": [] }, "members": { "names": ["jan"], "roles": [] } }'
```
上面的请求以anna的身份在mydatabase数据库下创建了_security文档。此后，这个数据库就不能匿名访问了，会报错。如：
```
curl http://10.0.7.105:5984/mydatabase/
```
而使用-u参数启用基础认证后就不报错了，如：
```
curl -u jan:orange http://10.0.7.105:5984/mydatabase/
```
将jan提升为mydatabase数据库管理员的请求：
```
curl -X PUT http://10.0.7.105:5984/mydatabase/_security \
     -u anna:secret \
     -H "Content-Type: application/json" \
     -d '{"admins": { "names": [jan], "roles": ["mydatabase_admin"] }, "members": { "names": [], "roles": [] } }'
```
## design文档的例子
```
{
"language": "javascript",

"validate_doc_update": "function(newDoc, oldDoc, userCtx) {
    function require(field, message) {
        message = message || "Document must have a " + field;
        if (!newDoc[field]) throw({forbidden : message});
    };
    function unchanged(field) {
        if (oldDoc && toJSON(oldDoc[field] != toJSON(newDoc[field])) {
        }
    };

    if (newDoc.type == "fortune") {
        require("body");
        require("sequence_id");
        require("created_at");
        unchanged("sequence_id");
        unchanged("created_at");
    }
};",

"views": { 
    "fortune_count": {
        "map": "function(doc) { if(doc.sequence_id && doc.body) { emit(doc.sequence_id, doc.body); }",
        "reduce": "_count"
        }
},

"shows": {
}
}
```
