# 说明

scala 官方 tutorial 第一章中重点介绍了集合类的特性：

* 通用性
* 一致性
* 功能丰富性

类似于 clojure 等设计思路

* 对集合作出抽象数据结构，有丰富的通用的 API  和函数式的用法
* 使用高阶函数（并提供语法支持）
* 方便在 API 兼容的情况下切换到多处理器支持等不同实现数据结构

优于 clojure 的地方在于

* scala 是静态类型语言，通过函数参数显式声明，在编译时会检测出很多类型错误

# 可变与不可变

集合类存在三种变体：

* root: `scala.collection`
* mutable: `scala.collection.mutable`
* immutable: `scala.collection.immutable`

> 为了方便和向后兼容性，一些导入类型在包 `scala` 中有别名，如：`List`, `Traversable`, `Iterable`, `Seq`, `IndexedSeq`, `Iterator`, `Stream`, `Vector`, `StringBuilder`, `Range`

集合类的抽象层次依次为：

1. Traverable
2. Iterable
3. Seq, Set, Map
4. ...

Traversable 类提供了所有集合支持的 API，例如，Traversable 类的 map 方法会返回另一个 Traversable 对象作为结果，但是这个结果类型在子类中被重写了。例如，在一个 List 上调用 map 会生成一个  List, 在 Set 上调用会再生成一个 Set，这种行为称为 __返回类型一致性原则__
