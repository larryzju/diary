# 版本

以版本 `a331c106f6ce4caa22f859e453cd4774cb9e2d25` 为例说明

# IDT

IDT ( Interrupt Descriptor Table ) 指定了一个数组（最大有 256 个），数组元素的类型为 `struct idt_entry` （表示 80386 INTERRUPT GATE 数据类型）。

IDT 数组通过 `struct idt_ptr` 结构来指向，其中包括：起始位置 `base` 和长度 `limit`

> TODO:
> 为什么 idtp.limit 保存的大小是 `(sizeof(struct idt_entry) * 256 ) - 1` 

## 流程


1. `main()`  在 main.c
2. `idt_install()` 在 idt.c
3. `idt_load()` 在 start.asm 52 行
4. 调用 `lidt` 指令加载指针 `struct idt_ptr`

# Exception Handler

IDT 表中 256 个中断号分为两部分：

1. 前 32 个为系统保留 （NMI 中断和异常）
2. 用户自定义外部中断（通过 8259A 可编程中断控制器）在 [32,256)范围

系统启动时将初始化前 32 个 NMI 中断处理过程（在 `isrs_install() 函数` 中）注册到相应的 IDT 注册号中，关键有以下两个值：

1. offset：指向处理的回调位置
2. selector：选择 GDT 或 LDT 表中的指定 segment 位置

## 流程

1. `main()` 在 main.c
2. `isrs_install()` 在 isrs.c
3. `idt_set_gate()` 在 isrs.c，将 `_isr[0-32]` 中的定义的位置设置到相应 IDT 中断号的 offset，选择 GDT 中的第一号表作为 segment 表（0x08），参见 [segment selector](segment.md#selector)

当有异常发生时，系统会进行中断处理:

1. 查找中断表，调用相应的 `_isr[0-32]` 代码
2. `_isr[0-32]` 压栈相应的 err_code，并调用 `isr_common_stub`
3. `isr_common_stub` 压栈各种寄存器，并调用 `fault_handler`（在isrs.c定义）
4. `fault_handler` 查找中断号，并打印文本信息


## 问题

* `struct regs` 中的最后一组内容信息 `eip`, `cs`, `eflags`, `useresp`, `ss` 是何时被压栈的？？？

* 代码中 `_isr8`, `_isr10`, `_isr11`, `_isr12`, `_isr13`, `_isr14` 等实现中没有压栈 `err_code` 是否 BUG?

