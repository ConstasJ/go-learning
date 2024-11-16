# Go的map
和Java的Map<T>不同，Go的Map并不是一个对象，而是一个数据类型，可以通过`make()`创建。在这一点上有点类似于JavaScript的`Object`和Python的`dict`。  
不过之前看有人在质疑这么做的必要性，以及这东西的**线程安全性**？后续可以关注一下。