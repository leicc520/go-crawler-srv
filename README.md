# go-crawler-srv

main.go        程序入口
spider.srv     通用的爬虫任务解析器整合
parser.srv     解析器处理逻辑主程序整合    
dispatcher.srv 任务调度器逻辑主程序整合
lib            第三方类库
   --sls       阿里云sls日志
   --proxy     代理服务
   --oss       oss存储落盘适配器
   --db        数据库落盘适配器
   --queue     队列处理逻辑


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