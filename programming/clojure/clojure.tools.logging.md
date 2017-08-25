# 提要

clojure `logging` 宏封装不同 log 库的实现细节，如 slf4j, Apache commons-logging, log4j 以及 java.util.logging

本日记主要目标是大概弄清楚 java 系中的各个库的关系，特别是 slf4j 和 log4j 的区别。弄清楚 clojure 的 logging 的使用和配置方法


# 基本概念

## log 级别

自下而上依次为：

1. trace
2. debug 
3. info
4. warn
5. error
6. fatal

## 基本操作

log 实现应该有以下方法：

* 各个级别日志是否启用 ( enabled? )
* 各个级别有独立的 write! 接口 

# 分类

查看 `clojure/tools/logging/impl.clj` 代码，发现四种 log 库 API 上两两分类：

* slf4j 和 cl (org.apache.commons.logging.Log) 相似
* log4j 和 jul (java.util.logging.Logger) 相似

从封装角度来看，感觉 slf4j 接口比较友好

# 对比

## slf4j vs log4j

引用自 [stackexchange](http://softwareengineering.stackexchange.com/questions/108683/slf4j-vs-log4j-which-one-to-prefer)

> SLF4J is basically an abstraction layer. It is not a logging implementation. It means that if you're writing a library and you use SLF4J, you can give that library to someone else to use and they can choose which logging implementation to use with SLF4J, e.g. log4j or the Java logging API. It helps prevent projects from being dependent on lots of loggin APIs just because they use libraries that are dependent on them.
>
> So, to summarise: SLF4J does not replace log4j, they work together. It removes the dependency on lo4j from your library/app.

简单的总结，log4j 是具体的实现，slf4j 是在多种实现上的封装接口。两者都可以直接用，但 log4j 同时关注了具体的日志实现

回头再看 [SLF4J 主页](http://www.slf4j.org/) 说明，引用如下：

> The Simple Logging Facade for Java (SLF4J) serves as a simple facade or abstraction for various logging frameworks ( e.g. java.util.logging, logback, log4j ) allowing the end user to plug in the desired logging framework at *deployment time.*

大概是 Java 系列有太多 log 实现，slf4j 定义了一套 API 让程序员与具体日志管理分离。Linux syslog 是否也是一样的结构（将实现与接口分开）？？

# slf4j vs cl

参考 [cl 官网](http://commons.apache.org/proper/commons-logging/)

> The Logging package is an ultra-thin bridge between different logging implementations. A library that uses the commons-logging API can be used with any logging implementatino at **runtime**. Commons-logging comes with support for a number of popular logging implementations, and writing adapters for others is a reasonably simple task.
>
> ... Using commons-logging does allow the application to change to a different logging implementation without recompiling code.

cl 与 slf4j 都是对不同实现的封装，两者的不同参考 [Simplifying the distinction between SLF4J and commons logging](http://jayunit100.blogspot.com/2013/10/simplifying-distinction-between-sl4j.html) 一文

![slf4j vs cl](https://chart.googleapis.com/chart?chl=+digraph+slf4j+%7B%0D%0A+++++slf4j+-%3E+slf4jbinding+%5Blabel%3D%22looks+for%22%5D+%3B%0D%0A+++++slf4jbinding+-%3E+implementation+%5Blabel%3D%22directly+translates+slf4j+calls+to%22%5D+%3B+%0D%0A%0D%0A+++++commonslogging+-%3E+magic_layer+%5Blabel%3D%22uses+its%22%5D+%3B%0D%0A+++++magic_layer+-%3E+%22org.apache.commons.logging.Log%22+%5Blabel%3D%22looks+in+.properties+or+system+properties%22%5D%0D%0A+++++magic_layer++-%3E+log4j+%5Blabel%3D%22nvm+try+this!%22%5D%3B%0D%0A+++++magic_layer++-%3E+JDKLogger+%5Blabel%3D%22or+this%3F%22%5D%3B+%0D%0A+++++magic_layer++-%3E+simples+%5Blabel%3D%22all+else+fails...%22%5D+%3B+%0D%0A%0D%0A%0D%0A+%7D%0D%0A++++++++&cht=gv)

粗略地看了一下，slf4j 实现更简洁、效率更高，blabla

# logback vs log4j

logback 重新实现了 log4j，效率更高，有原生的 SLF4J binder。所以还是用 slf4j 吧

# 如何不添加 log4j.properties 而动态添加日志规则

首先约定 clojure 使用的 `*logger-factory*` 是 slf4j，而 slf4j 封装了 log4j 接口，因此日志规则需要用 log4j 的。

网上查找 "log4j properties dynamic"，解决方法是构造一个 Properties 对象，然后使用 `org.apache.log4j.PropertyConfigurator` 加载之。暂记到这里，需要回头研究下 log4j 的配置语法，再来验证
