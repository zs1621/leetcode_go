

## 符号表定义

 > 符合表是一种存储 键值对的数据结构, 
>支持两种操作: 插入(Put), 将将一组新的键值对存入表中;查找(Get), 根据给定的键得到相应的值

 - SequentialSearchST: 无序列表
 - BinarySearchST: 有序数组
 - BST: 二叉搜索树
 

### 运行时间对比

|结构\测试文件 | leize1m.txt | tale.txt |
| ----- | ----- | ----- |
| SequentialSearchST | - | 2.559444795s |
| BinarySearchST | 4m6.318207461s | 239.376313ms |
| BST | 22.449s | 182.501156ms |