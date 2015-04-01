##CASE中医病历库
部署方案

###节点软件
####文件夹结构
<pre>
.
├── static
|   ├── css
|   ├── image
|   ├── js
|   └── template
├── data  
|   ├── person.yaml
|   ├── watchlist.yaml
|   ├── sharedata
|   |   ├── xxx.person.yaml
|   |   ├── xxx.key.pub
|   |   ├── person.yyy.com.yaml
|   |   ├── person.yyy.cod.yaml
|   |   ├── person.yyy.pool.yaml
|   |   └── ...
|   ├── index.yaml
|   ├── contract
|   |   ├── 1.yaml
|   |   ├── 2.yaml
|   |   └── ...
|   ├── ticket
|   |   ├── 1.yaml
|   |   ├── 1.budget.pool.aaa.yaml
|   |   ├── 2.yaml
|   |   └── ...
|   ├── 20150101010203
|   |   ├── baseline.yaml
|   |   ├── index.yaml
|   |   └── ...
|   ├── YYMMDDhhmmss
|   |   ├── baseline.yaml
|   |   ├── index.yaml
|   |   └── ...
|   └── ...
├── client.v0.1.exe
├── client.v0.2.exe
├── client.v....exe
├── config.yaml
├── pointer.yaml
├── key.pub
├── key.sec
└── README.md
</pre>