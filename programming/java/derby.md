# derby tutorial 备忘

## 安装与验证

1. 解压
2. 设置 `DERBY_INSTALL` 环境变量，并使用 `./bin/setXXXX` 来设置 CLASSPATH
3. 或者手动设置 CLASSPATH
   - derby.jar: 基本的 embedded 服务
   - derbytools.jar: ij, sysinfo 工具
   - derbynet.jar: Network Server
   - derbyrun.jar: 服务管理
4. 使用 `java org.apache.derby.tools.sysinfo` 来验证

## ij 命令

> 类似于 beeline 之于 hive

通过 `java org.apache.derby.tools.ij` 进入

* `connect 'jdbc:derby:MyDbTest;create=true'`
* `run 'my_file.sql'`

## embedded vs network server

* 嵌入模式下 derby 实例与 JDBC 运行在同一个 JVM 中，同一时刻只能有一个 JVM 操作库
* network server 通过框架启动 derby 实例，并以 JDBC ( network client ) 方式处理多客户端请求

## 启动 network server

* `bin/startNetworkServer` 启动监听 1527 端口
* `java -jar derbyrun.jar server start`

