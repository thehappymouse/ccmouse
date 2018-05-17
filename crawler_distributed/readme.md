## 根据并发爬虫，开展分布式改造

使用jsonrpc,分离 saver, worker, 去重等节点

第一步：

saver分离

运行步骤：
启动 persist/server/main.go

再启动 main.go

解析器的RPC, 难点在于解析器函数的序列化和反序列化

把 Work里的 parserfunc转换成 Parsr接口


运行步骤：

1，启 saver
2, 启多个worker    worker_main --port=9002    worker_main --port=9003
3，启引擎 distributeed --hosts=":9002,:9003"
