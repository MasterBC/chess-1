# room(逻辑)

## 设计理念
游戏服务器对agent只提供一个接口， 即:

> rpc Stream(stream Room.Frame) returns (stream Room.Frame);

接收来自agent的请求Frame流，并返回给agent响应Frame流

来自设备的数据包，通过agent后直接透传到room server, Frame大体分为两类：  

1. 链路控制（register, kick)     
2. 来自设备的经过agent解密后的数据包 (message)       

数据包(message)格式为:      

> 协议号＋数据

        +----------------------------------+     
        | PROTO(2) | PAYLOAD(n)            |     
        +----------------------------------+     

在client_handler目录中绑定对应函数进行处理

## 安装
参考Dockerfile
