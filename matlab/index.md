## 矩阵

### 构造

```matlab
A = [1 2 3; 4 5 6];  # 两行，三列
B = rand(5,5);       # 5x5 随机
C = magic(3);        # 行，列相加均相等
D = [A; C];		     # 上下拼接
```

### 索引

1. 通过下标: `A(2,3)` 第二行，第三列
2. 单下标：`A(5)` 第一行，第三列
3. 范围
   1. `A(1,:)` 第一行
   2. `B(2:4, 2:4)` 

### 环境

1. `whos` 列出变量（类似 env in shell）
2. `save myfile.mat` 保存环境（运行上下文）
3. `load myfile.mat` 加载上下文
4. `clear` 清除环境

### 字符串

- 用单引号引用
- `'he''s Jerry'` 用单引号转义单引号
- 字符串是行向量，可以拼接: `[a, '-', b]`
- 用 `num2str` 或 `int2str` 将数字转换为字符串