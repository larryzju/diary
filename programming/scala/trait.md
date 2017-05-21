# 概述

trait 类似 java 的 interface，用于定义类的方法模板。


# 继承(inherit)与混成(mixin)


mixin 类比于电影中合成技术（混成），对于两个相同规格的场景（基础类），混成出一个全新的场景（extends ... with）。

trait 可以扩展(extends) 抽象类，实现扩展的方法。并在同样继承(inherit) 抽象类的类中，混成 trait 扩展的方法。

优点在于，将扩展的通用实现加入到朴素的原始数据结构中。

如：

```scala
abstract class AbsIterator {
  type T
  def hasNext: Boolean
  def next: T
}

trait RichIterator extends AbsIterator {
  def foreach( f: T => Unit ) { while ( hasNext ) f(next) }
}

class StringIterator( s: String ) extends AbsIterator {
  type T = Char
  private var i = 0
  def hasNext = i < s.length()
  def next  = { val ch = s charAt i; i += 1; ch }
}

object StringIteratorTest {
  def main( args: Array[String] ) {
    class Iter extends StringIterator( args(0) ) with RichIterator
    val iter = new Iter
    iter foreach println
  }
}
```


# 注意事项

* trait 不能有构造函数
* 定义的函数可以是抽象的，也可以是具体的
* 具体的函数可以依赖于抽象函数
