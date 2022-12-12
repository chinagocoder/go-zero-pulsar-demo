# go-zero-pulsar-demo

> go-zero 使用pulsar消息队列的demo，pulsar框架，引用https://github.com/zeromicro/go-queue

根据go-zero目录规范，go-zero的业务拆分大概分为以下几种功能模块,因此demo的生产者在api模块，消费者在rmq模块
```
xxx
├── api //  http访问服务，业务需求实现
├── cronjob // 定时任务，定时数据更新业务
├── rmq // 消息处理系统：mq和dq，处理一些高并发和延时消息业务
├── rpc // rpc服务，给其他子系统提供基础数据访问
└── script // 脚本，处理一些临时运营需求，临时数据修复
```
