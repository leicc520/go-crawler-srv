# go-crawler-srv

main.go     程序入口
adapter     存放各种适配器
   --oss    oss存储落盘适配器
   --notify 第三方通知
   --db     数据库落盘适配器
spider      通用的爬虫任务整合
lib         第三方类库
   --sls    阿里云sls日志
   --proxy  代理服务



第三方插件处理逻辑
https://chromedriver.chromium.org/downloads

RabbitMQ
yum install rabbitmq-server
rabbitmq-plugins enable rabbitmq_management
http://ip:15672
密码 guest | guest