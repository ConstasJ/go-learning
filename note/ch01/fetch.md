# Go中的HTTP请求
看来Go标准库中似乎也有http请求库呢。目前看来基本够用了，不过如果涉及到CORS等高级功能的话，可能还是需要引入第三方库。  
然后还是那个古老的问题——`if err != nil`……  
涉及到Http请求的时候起码三个这玩意，看起来真糟心……

## http.Get()
注意一下，和新ES的`fetch()`类似，`http.Get()`的`Body()`是一个IO流。  
换言之，你需要把它当作流来处理，并记得关闭！  

幸好Go的`io`包提供了不少有用的工具函数，这个过程不至于太困难。  
如果你只是要从响应体中读出数据，可以使用`ioutil.ReadAll()`，这个函数会帮你读取所有的数据并返回一个字节数组。  
当然，如果你的数据量很大，这个函数可能会导致内存溢出（缓冲区size过大的问题），所以最好还是使用`io.Copy()`来处理，尤其是你只是想把数据写到文件或标准输入输出流中的时候。

# 关于io.Copy()
这个函数的返回值是拷贝的字节数，一个int64。  
如果不需要知道这个数据（大部分情况我觉得都不太需要）可以使用之前提到的`_`来忽略。