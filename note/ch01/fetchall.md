# Go的协程机制
协程机制可以说是Go的一个核心竞争力了。其低开销和易于使用的协程使得其可以轻松支持高并发任务，这也是Go语言在网络编程和高并发编程中的优势所在。
而且，相比于Kotlin的协程，JS的async/await，Python的asyncio等，Go的协程机制不具备“传染性”，即可以无缝实现同步和异步的转换，而不需要额外的语法糖。
关于具体的用法，将会在后面具体学习到的时候再详细介绍。

Go的协程基于goroutine和channel实现。  
当需要进行一些非阻塞操作的时候，使用`go`关键字即可创建一个goroutine，在需要结果的地方使用Channel接收结果。  
换言之，Golang的协程机制是基于消息传递的，而不是共享内存的。这样可以避免很多并发问题，比如死锁、资源竞争等。  
不过也存在一个问题，似乎没法像JS那样使用`Promise.all()`这样的方法来等待多个协程的结果>_<