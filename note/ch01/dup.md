# Go的map
和Java的Map<T>不同，Go的Map并不是一个对象，而是一个复合数据类型，可以通过`make()`创建。在这一点上有点类似于JavaScript的`Object`和Python的`dict`。  
不过之前看有人在质疑这么做的必要性，以及这东西的**线程安全性**？后续可以关注一下。  

和其他语言的Map一样，Go的Map也是一个键值对的集合，其中键和值的类型可以是任意类型，但键的类型必须是可以用`==`比较的类型。  
（等等，Golang中的`==`是不是只能用于基本类型？还是说和Kotlin一样是`equals()`的语法糖？）  
不过说到这里，JS的`Object`是强制要求键必须是字符串的，不然会出错。`dict`则没有这个限制。  
Java的`HashMap`/`TreeMap`、Kotlin的`Map`/`MutableMap`也是可以用任意类型作为键的。  

文章中也提到了，Go的Map底层使用哈希表来实现，和Java的`HashMap`一致。  
似乎`dict`也是使用哈希表实现的。

具体到使用层面，Go的Map也可以使用索引运算符`[]`来访问元素，和JS的`Object`、Python的`dict`一样。  
Java的Map则需要调用`get()`方法。不过Java 8之后也支持使用`[]`来访问元素。  
Kotlin的`Map`/`MutableMap`也是使用`[]`来访问元素。

## Map的遍历
Go的Map可以使用`for range`来遍历，和Python的`dict`一样。  
不过有一点需要注意一下：Go的Map遍历时不保证顺序。如果需要确保有序，可能需要一点额外的工作。  
根据StackOverflow上的回答，这是有意为之的，确保程序不会依赖于特定的遍历顺序。
[相关StackOverflow问答](https://stackoverflow.com/questions/11853396/google-go-lang-assignment-order)
> A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type.  
> 翻译：Map是一种无序的元素集合，由另一种类型的唯一键索引，称为键类型。  
> -- [Go语言规范](https://golang.org/ref/spec#Map_types)  

这里的“无序”指的是遍历时的顺序，而不是键值对的存储顺序。

## Map作为参数传递
Go的Map是引用类型，所以在函数调用时，传递的是指向Map的指针。  
严格来说，这里的指针其实并不是原本的指针，而是原指针的副本。不过确实是指向同一个内存区域的。  
以及，这里的传递并不需要使用`&`来取地址，直接传递Map变量即可。

# 关于Scanner的终止输入（stdin）
这里有一点需要注意一下：如何中止`Scan()`的循环？  
其实不难：如果要结束输入，先输入一个空格，然后按`Ctrl+D`（Linux）或`Ctrl+Z`（Windows）即可。

# Go的fmt.Printf()（格式化输出）
看上去和C语言的`printf()`差不多。使用的格式化占位符也参考了C语言`stdio.h`的`printf()`。  
参数本身也和C语言的`printf()`类似，第一个参数是格式化字符串，后面的参数是要输出的内容。
之前使用的`fmt.Println()`是没有格式化功能的，只能输出字符串。  
这一点更接近于Java的`System.out.println()`，JS的`console.log()`，Python的`print()`。  
参考这一惯例，基本上以`f`结尾的函数都是格式化输出的，且不自带换行符。  
而以`ln`结尾的函数则是普通输出，自带换行符。  

# 关于Go的变量及其类型
Go是一种静态类型语言，变量的类型在编译时就已经确定。同样的，变量的类型也是不可变的。  
不过和之前见过的其他语言都不一样，Go变量的类型后置，中间却使用空格分隔变量名和类型。这和C和Java的变量类型前置，TypeScript和Kotlin的变量类型后置、使用冒号分隔有所不同。  
以及，Go似乎是存在指针类型的，不过功能相比于C和C++有所简化，比如说没有指针运算。  

顺带一提，map的类型是`map[keyType]valueType`。颇有意思。  
Java、Kotlin和TypeScript都是使用泛型来表示Map的键和值的类型。

# Go的错误处理机制
众所周知，Gopher的键盘需要一个键，用于打出以下代码段：
```go
if err != nil {
    // 错误处理流程
}
```
这里反应了Go的错误处理机制：Go的错误处理是通过返回值来实现的。  
上次见到使用返回值来处理错误的还是Node.js风格的回调函数:
```javascript
fs.readFile('file.txt', (err, data) => {
    if (err) {
        // 错误处理流程
    }
});
```
我个人其实不太喜欢这种方式。相对而言还是更习惯`try-catch`。起码把错误处理和正常的返回值分离开来吧……这也是我不太喜欢Go的一个地方。

# Go的包级函数
Go的包级函数是指在包内定义的函数，可以被包内的其他函数调用。  
和TypeScript，Kotlin一样，无需注意函数的先后顺序。不过出于可读性考虑，建议制定一定的规范。  
如果函数名以大写字母开头，则表示该函数是公有的，可以被其他包调用。