# 说明

使用 maven-archetype-plugin 建立一个新的工程

```
mvn archetype:generate
```

默认情况会从两个位置下载 XML 文件：

1. remote: https://repo.maven.apache.org/maven2/archetype-catalog.xml
2. local: 当前目录下的 archetype-catalog.xml

然后进入一个交互式选择模式，通过指定模板编号，填写相应的 GAV 名称，生成一个空的项目



# 示例

如果需要新建一个 scala 工程，需要以下步骤

1. `mvn archetype:generate` 进入交互式命令行（可能需要一段时间）
2. 输入 `scala` 来筛选 scala 相关的模板
3. 选择具体的编号，这里选择 `7: remote -> net.alchim31.maven:scala-archetype-simple`
4. 填写 GAV 信息，按 "Y" 生成工程


# 问题

aliyun 的 mirrors 会导致 remote 下载异常？

# 速度

xml 文件有 5M 大小，网络比较慢的情况下，下载时间过长。

可以先用 wget 下载文件到 `~/.m2/repository` 目录下，使用 `mvn  archetype:generate -DarchetypeCatalog=local` 加载之

另外，可以为 `maven -X` 参数打开 debug 来调试
