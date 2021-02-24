# dingtalk
http proxy from alertmanager message to dingding

# Run with docker
You can run application with docker.

```$xslt
docker build -t dingtalk .
```

then
```$xslt
docker run --name dingtalk -p 8060:8060 \&
-e "ACCESS_TOKEN=xxxx" -e "SERVER_PORT=8060" dingtalk
```

## environment params

| param | type | default |description | example |
| :--- | :--- | :--- | :--- | :--- |
| ACCESS_TOKEN | string | "" | dingding robot access token | xxxxxxxxx |
| SERVER_PORT | string | 8060 |server 启动后所占用端口 | 8060 |