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