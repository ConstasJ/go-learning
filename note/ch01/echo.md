# 关于变量初始化
Go语言使用var关键字来声明变量（ES6以前的JS也在用，问题不大），类型名后置。后置类型名称本身也算是现代语言的常规操作了（Kotlin和TypeScript也是这么干的），不过Go和它们不一样的是，类型名和变量名之间是用空格分隔的。  
而且，Go还支持短变量声明，即使用`:=`来声明变量并初始化。这种方式只能用在函数内部，不能用于包级别的变量声明。
# 关于For循环
Go语言的`for`循环有以下几种格式：  
1. 完整的C样式的`for`循环：
   ```go
    for init; condition; post {
        // 循环体
    }
   ```
   其中`init`、`condition`和`post`都是可选的。  
   换言之，可以通过只使用`condition`模拟while循环，以及通过省略所有部分模拟while(true)无限循环。
2. 类似Java中的`foreach`循环格式：
   ```go
    for index, value := range array {
        // 循环体
    }
   ```
   其中`index`和`value`是可选的，`index`是数组的索引，`value`是数组元素的值。  
   如果不需要`index`，可以使用`_`替代。  
   或许这里的`array`不一定非得是个数组？我觉得应该但凡是个可遍历对象都可以吧。