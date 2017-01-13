# 概述

## inline

inline 用于内联插入代码块，是一种代码优化策略。减少函数调用造成的开销

## asm

关键字 `asm` （也可以用 `__asm__`） 用于在 C 代码中插入汇编代码

## 语法

使用 AT&T 汇编语法，分为基本语法和扩展语法。

扩展语法主要原因是：

* 需要传入 C 的变量
* 需要告诉 GCC 指令修改了些什么，以便 GCC 进行优化




## volatile

用于阻止 GCC 优化代码，避免编译器将这部分代码移动或删除

# AT&T 语法

1. Source-Destination Ordering: `opcode src dst`
2. Register Nameing: `%eax`
3. Immediate Operand: `$0x08`
4. Operand Size: `movl`
5. Memory Operands: `disp(base,index,scale)`

# ASM 语法

	asm ( assembler template
		: output operands
		: input  operands
		: list of clobbered registers );
		
其中输入和输出的语法为逗号分隔的变量描述，变量描述如下：

	"c" (count)
	
前者称为 constraint，用于标记使用的类型（寄存器？内存？立即数）及修饰词

后者为 C 的变量或字面常量

示例如下：

	asm ( "movl %1, %%eax;
	       movl %%eax, %0;"
	    :  "=r" (b)
		:  "r"  (a)
		:  "%eax" );

第三项称为 clobber list，用于告知 GCC 哪些寄存器被修改，以便进行保护。**不需要列入输入输出寄存器**，因为 GCC 知道它自己修改了这些寄存器中的哪些

# 参考资料

[GCC inline assembly howto](http://www.ibiblio.org/gferg/ldp/GCC-Inline-Assembly-HOWTO.html)
