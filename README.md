# 简单的服务器框架

使用类似于http框架，
设置router，然后只需要关心router之间的逻辑，不需要关心对于连接的处理

内部使用了非阻塞的通道来接发数据

里面包含了ini类似的配置文件，很多参数可以默认为空，为空则使用框架自己的配置

心跳包默认已经为使用者实现，但是心跳数据需要使用者手动去注入，每次心跳都会调用注入的心跳方法，所以可以保证心跳每次传递的，都是服务器想要发送给客户端的数据，心跳默认使用时间轮的算法去实现，目前没有其他的算法，时间轮的测试，可以看libtime的包。
