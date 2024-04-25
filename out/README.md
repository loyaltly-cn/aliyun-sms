# 阿里云短信服务



## 启动

> 运行 sms.sh / sms.exe 会自动读取config.json的内容 
>
> shell 会根据json内配置的端口启动http服务

## 注意事项

- 使用前请确保已获取config.json对应的参数
- 修改config.json参数时请勿修改任何的key
- sms.sh 只能在amd64下的linux环境运行
- 请确保config.json 在sms.sh / sms.exe同级目录下
- 配置http端口时请确保port的值为string类型

### API

> 成功启动通过 ip:port/uri 访问 API

| uri  | method | param | response | desc |
|------| ------ | ----- | -------- |------|
| test | GET    | /     | "test"   | 测试 |
| sms  | POST   | code:strring <br/> phone:string | true | 发送验证码|

> Content-Type: application/json
