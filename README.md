# go-zero-pulsar-demo

> go-zero 使用pulsar消息队列的demo，pulsar框架，引用https://github.com/chinagocoder/go-queue
# 目录说明
根据go-zero目录规范，go-zero的业务拆分大概分为以下几种功能模块
```
xxx
├── api //  http访问服务，业务需求实现
├── cronjob // 定时任务，定时数据更新业务
├── rmq // 消息处理系统：mq和dq，处理一些高并发和延时消息业务
├── rpc // rpc服务，给其他子系统提供基础数据访问
└── script // 脚本，处理一些临时运营需求，临时数据修复
```
因此demo的目录拆分如下
```
demo
├── api // pulsar生产者
├── rmq // pulsar消费者
```
# 运行测试
## 修改pulsar配置
```yaml
Pulsar:
  Brokers:
    - 127.0.0.1:6650
  Topic: topic_demo
  Conns: 1
  Processors: 1
  SubscriptionName: sub_topic_demo
  # 如果不需要认证，请删除AuthName和AuthParams
  AuthName: token
  AuthParams: '{"token":"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJhZG1pbiJ9.CNYN8h04Z_wJcvNssVhcyZKDlqvwOSxmkXeOy6WH8pM"}'
```
## 生产者测试
```shell
#启动生产者
cd demo/api
go run demo.go -f etc/demo.yaml
#浏览器访问,参数msg为消息内容
http://127.0.0.1:8888/demo/produce?msg=hello
```
## 消息者测试
```shell
#启动消费者
cd demo/rmq
go run demo.go -f etc/demo.yaml
```