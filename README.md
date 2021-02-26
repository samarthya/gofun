# Learning communication
There are two basic constructs for concurrent programmin in GoLang.
- goroutines
- channels

In this space I will try and write examples exploring capabilities for both.

> Parallelism is the ability to make things run quickly by using multiple processors simultaneously. So concurrent programs may or may not be parallel.

#### You can read my blog on [concurrency](https://blog.samarthya.me/wps/2020/01/03/concurrency-golang/) & [race condition](https://blog.samarthya.me/wps/2020/10/06/race-condition/).

- The part of the application process that run concurrently are called goroutines.
- goroutine & a thread
- goroutines run in the same address space
- Lightweight: They are created with a 4KB memory stack-space on the heap.
- The stack of a goroutine grows and shrinks as needed.
- When the goroutine finishes silently, implying nothing is returned to the function which started it. 

# Channel
Go has a special type channel that takes care of communicating between goroutines.

```
var myChannel chan int
myChannel = make(chan int)

// Once done you need to close it
close(myChannel)
```

- Channel can be made of any type.
- Channel can transmit data of that type only.
- It is like FIFO data structure
- It is reference data type