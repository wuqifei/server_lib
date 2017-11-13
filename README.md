# 简单的TCP服务框架以及一些常用的库

测试代码，可以参照example里面的文件，性能测试，做了一些简单的测试，每个lib包，有些也有test类，可以参看

libnet库主要是一个tcp库，使用类似于http框架的使用
设置router，然后只需要关心router之间的逻辑，不需要关心对于连接的处理

内部使用了非阻塞的通道来接发数据

里面包含了ini类似的配置文件的库，可以用这个库，来配置自己的配置文件，详情可以参考libconf文件夹

去除代码中的用户管理模块，时间轮模块，保持代码功能单一性

如果只是单纯使用，tcp的库，则只需要引入libnet,libio,concurrent三个包,其他的包可以不引入

libtime，默认为用户实现了时间轮算法，可以直接使用
