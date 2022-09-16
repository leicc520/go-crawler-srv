# go-crawler-srv
```
本地启动的话需要配置一下环境变量
DCENV=loc
或者启动命令配置 -dcEnv=loc 也可以
启动目录配置到项目目录路径，让他能找到配置config目录加载配置

main.go        程序入口
api            http服务请求入口
adapter
   --command   所有命令行可以执行程序的注册器
   --spider    通用的爬虫任务解析器整合
   --wipo      品牌抓取的脚本处理逻辑
lib            第三方类库
   --proxy     代理服务
   --oss       oss存储落盘适配器
   --parser    采集解析器引擎处理逻辑
   --queue     队列处理逻辑
   sls.go      阿里云sls日志
   log.go      日志记录，有些日志只记录一次
   redis.go    redis全局的初始化
   config.go   配置管理,应用全局配置管理
plugins        第三方插件管理
service        第三方服务相关
   --dbsrv     数据库模型-自动生成即可
   --oss       OSS上传管理相关 
```

第三方插件处理逻辑
https://chromedriver.chromium.org/downloads
```
RabbitMQ
yum install rabbitmq-server
rabbitmq-plugins enable rabbitmq_management
http://ip:15672
密码 guest | guest

docker pull rabbitmq:3.10.7-management
docker run -d --name rabbitmq3.10.7-management -p 5672:5672 -p 15672:15672 -v `pwd`/data:/var/lib/rabbitmq --hostname rabbitmqMY -e RABBITMQ_DEFAULT_VHOST=spider -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=spider8888 --restart=always rabbitmq:3.10.7-management
```