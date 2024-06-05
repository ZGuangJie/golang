#### 1、概念辨析

##### 1.1 程序（program） 

​		是为完成特定任务、用某种语言编写的一组指令的集合，是一段静态的代码。(程序是**==静态==**的）

##### 1.2 进程（process）

​		是程序的一次**执行**过程。正在运行的一个程序，进程作为资源分配的单位，在内存中会为每个进程分配不同的内存区域。(进程是**==动态==**的)是一个动的过程，进程的生命周期：有它自身的**产生**、**存在**和**消亡**的过程。

##### 1.3 线程（thread）

​		**进程**可进一步细化为**线程**，是一个程序内部的一条**执行路径**。若一个进程同一时间并行执行多个线程，就是支持多线程的。

##### 1.4 协程（goroutine）

​		又称为**微线程**，纤程，协程是一种用户态的**轻量级线程**作用：在执行A函数的时候，可以随时中断，去执行B函数，然后中断继续执行A函数（可以自动切换），注意这一切换过程并不是函数调用（没有调用语句)，过程很像多线程，然而协程中只有一个线程在执行（协程的本质是个**单线程**）。

​		对于单线程下，我们不可避免程序中出现**I/O**操作，但如果我们能在自己的程序中（即用户程序级别，而非操作系统级别）控制单线程下的多个任务能在个任务遇到阻塞时就将寄存器上下文和栈保存到某个其他地方，然后切换到另外一个任务去计算。在任务切回来的时候，恢复先前保存的寄存器上下文和栈，这样就保证了该线程能够最大限度地处于就绪态，即随时都可以被cpu执行的状态，相当于我们在用户程序级别将自己的I/O操作最大限度地隐藏起来，从而可以迷惑操作系统，让其看到：该线程好像是一直在计算，I/O比较少，从而会更多的将cpu的执行权限分配给我们的线程（注意**：线程是CPU控制**的，而**协程是程序自身控制**的，属于程序级别的切换，操作系统完全感知不到，因而更加轻量级)。



#### 2、协程的使用

​		golang启动协程十分简单，协 程实质上是一个函数。直接go functionName()即可。

```go
package main

import (
	"fmt"
	"time"
)
func test() {
	for i := 1; i < 5; i++ {
		fmt.Println("I'm gorotine...")
		time.Sleep(time.Second)
	}
}
func main() {
	// 使用关键字 go 启动一个协程
	go test()
	// 主线程与协程交替运行
	for i := 1; i < 5; i++ {
		fmt.Println("I'm threading...")
		time.Sleep(time.Second)
	}
}

```

##### 2.1 多协程启动

​		协程可以同时启动多个，而且协程之间切换不浪费CPU时间，非常高效。下面是使用for循环启动多个协程。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 5; i++ {
		// 使用一个循环启动多个协程
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Second)
}
/*
output:
	乱序的1、2、3、4
// 这个结果说明协程启动很快，而且刚开始启动时并未执行
*/
```

##### 2.2 互斥锁

​		Golang使用sync库来解决协程之间**访问公共资源出现的竞争问题**。下面add()和sub()函数在访问同一变量number时，如果不加锁，则会导致，加n次后在减n次结果不为0，这是因为在协程执行过程中，随时发生切换，导致number变量变换不同步。使用sync下的.Mutex加锁解决这个问题。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)
var numer int
var lock sync.Mutex

func add() {
	for i := 0; i < 10000; i++ {
		// 访问公共变量时加锁
		lock.Lock()
		numer += 1
		// 访问结束时解锁
		lock.Unlock()
	}
}
func sub() {
	for i := 0; i < 10000; i++ {
		// 访问公共变量时加锁
		lock.Lock()
		numer -= 1
		// 访问结束时解锁
		lock.Unlock()
	}
}
func main() {
	go add()
	go sub()
	time.Sleep(time.Second)
	fmt.Println(numer)
}
/*
output:
	理论上: -10000~10000之间的数都有可能
	// 加锁之后结果才为0
*/
```

##### 2.3 读写锁

​		虽然sync.Mutex能解决多协程访问同一变量的带来的竞争问题，但sync.Mutex锁在使用过程中只能有一个协程能拿到锁，这样就**严重影响了程序执行效率（需要等解锁）**。尤其是有些协程之间同时访问某一变量对该变量并无影响，因此只要给部分有影响的协程加锁即可。这就引出了**读写锁**。在golang中读写锁使用sync.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.RWMutex

func read() {
	// 调用读锁，多个读之间不影响，写的时候才影响
	lock.RLock()
	fmt.Println("我在读呢")
	time.Sleep(time.Second)
	fmt.Println("读取完毕")
	lock.RUnlock()
}
func writer() {
	// 调用写锁，同时只有一个协程拥有写锁
	lock.Lock()
	fmt.Println("在写")
	time.Sleep(time.Second * 3)
	fmt.Println("写完了")
	lock.Unlock()
}
func main() {
	for i := 0; i < 5; i++ {
		go read()
	}
	go writer()
	time.Sleep(time.Second * 10)
}
/*
 多个读协程都可以同时进入，在读协程完后写 协程 拿到锁之后别的协程都进不去
*/
```

​		**`一句话读多写少用读写锁，不知道谁多谁少用互斥锁。`**

##### 2.4 主死从随

​		一句话，协程是附属于某一个线程的，当线程结束以后，协程自动结束。可以理解为，协程就是一个特殊的函数。

```go
package main
import (
	"fmt"
)
func main() {
	// 使用关键字 go 启动一个协程
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("I'm gorotine...")
		}()
	}

	fmt.Println("I'm threading...")
}
/*
	按理说上面的程序会输出一个I'm threading...，5个I'm gorotine...，但是因为主线程在输出后就运行结束了，上面的协程也会一起结束，所以并不会输出5个I'm gorotine...。
*/
```

##### 2.5 使用WaitGroup控制协程退出

​		因为主死从随的缘故，在主线程结束后即使协程没有运行完也会随之结束。上面的一些程序为了解决这个问题，直接在主线程后面加了time.Sleep(time.Second)阻塞用来等待协程的运行完毕。但这太不优雅了.....Go琅提供了WaitGroup来解决这个问题，主要使用Add，Done，Wait三个方法，下面来看一下怎么使用：

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 使用关键字 go 启动一个协程
	for i := 0; i < 5; i++ {
		// 进程启动时，增加一个
		wg.Add(1)
		go func() {
			// 进程结束时，减少一个
			defer wg.Done()
			fmt.Println("I'm gorotine...")
		}()
	}
	fmt.Println("I'm threading...")
    // 调用wg.Wait()，等待所有gorotine执行完后才结束主线程
	wg.Wait()
}
/*
	使用 wg.Add(n)可增加n个协程，使用wg.Done()表示一个协程已经执行完毕。
	最后调用wg.Wait()等待所有协程执行完毕，在所有协程执行完毕以前，不会执行wg.Wait()后面的语句
*/
```

#### 3、 管道

​		
