# fabrik
A configuration center.

# 流程
每一个服务注册时向Fabrik发送注册请求，销毁时发送销毁请求。  
Fabrik存在心跳机制向每个服务发送Ping，检测已注册服务的健康状态。