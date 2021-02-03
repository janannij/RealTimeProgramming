Exercise 1 - Theory questions
-----------------------------
 
 ### What is the difference between concurrency and parallelism?
 > Concurrency is used about an appication that processes more than one task at the time, while parallelism is used about applications where the tasks are divided into sub-tasks, which are processed parallel.
 
 ### Why have machines become increasingly multicore in the past decade?
 > Because multicore CPUs perform much better than a single-core CPU at the same time. Multicore allows the machine to run multiple processes at the same time, which increases the performance when multitasking. 
 
 ### Why do we divide software (programs) into concurrently executing parts (like threads or processes)?
 (Or phrased differently: What problems do concurrency help in solving?)
 > As we divide the program into several, independent parts, each part can be executed concurrently. As we can execute the the parts individually, without having to wait for all of them, the program is much more effective this way.
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > Creating concurrent programs has both pros and cons. By dividing the program into concurrent parts, it is easier to develope the program and make it effective. But some errors may show up when executing the program, as the concurrent parts may interract with each other.
 
 ### What is the conceptual difference between threads and processes?
 > A process is any program in execution, while a thread is a segment of an process. A process can have multiple threads. A thread takes less time to terminate compared to a process, but on the other hand, it does not share memory with any other process, as a thread would do.
 
 ### Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they?
 > Fibers are considered to be lightweight, cooperative threads. A thread can consist of several fibers. The difference is that a thread may be interrupted at any time, while fibers switches between tasks as the user orders it. Fibers start and stop in defined places, so data integrity won't be much of an issue. 
 
 ### What is the Go-language's "goroutine"? A C/POSIX "pthread"?
 > The goroutine is a lightweight thread, which means it uses less resources compared to a regular thread. A pthread is multithreading used in C/POSIX.
 
 ### In Go, what does `func GOMAXPROCS(n int) int` change? 
 > This code changes the max. possible number of CPUs that can execute Go code at the same time, as well as returning the original number.



 
 
 
 
