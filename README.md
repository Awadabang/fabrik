# Fabrik

A configuration center.

# 流程

每个服务在构建时生成一个AccessKey，服务注册时向Fabrik发送注册请求，携带AccessKey。  
每个服务销毁时发送销毁请求，携带AccessKey。  
Fabrik存在心跳机制向每个服务发送Ping，检测已注册服务的健康状态，取消维护不健康的服务。  
主动加被动双重保证服务的状态维护。    
Ymonitor和Worker注册后，Fabrik保存各自ServiceName、ServiceURL、AccessKey  
Ymonitor端MongoDB监听Alert Autos字段的状态，如果Length>0，就向Fabrik发送该Alert的Id和Autos  

# 项目目标
实现gRPC注册中心，服务发现，服务注册，心跳