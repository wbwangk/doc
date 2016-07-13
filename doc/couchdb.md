---
layout: page
title: 免费https证书(Let's Encrypt)
---

## curl与couchdb基础认证
1. curl --user anna:secret  https://couchdb.imaicloud.com/_config
2. curl https://anna:secret@couchdb.imaicloud.com/_config
以上两种url用法证明都可以。第2种写法在浏览器地址栏中输入是不行的。

创建一个叫somedatabase的数据库：curl -X PUT --user anna:secret https://couchdb.imaicloud.com/somedatabase
