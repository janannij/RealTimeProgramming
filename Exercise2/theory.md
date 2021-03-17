Exercise 2 - Theory questions
-----------------------------

### What is an atomic operation?
> In concurrent programming, atomic operations are operations that cannot be interrupted, the operation must be fully executed before any other processes can be executed. This means that atomic operations are isolated from other operations.

### What is a critical section?
> A critical section is a protected section in a program that can only be executed by one process at the time.

### What is the difference between race conditions and data races?
> A data race occurs when two or more threads try to access the same variable simultaneously, at least one of the threads writes to it and there is no synchronization. Race condition is a semantic error that occurs in the timing or ordering of events. A race condition may be a result of data race.

### What are the differences between semaphores, binary semaphores, and mutexes?
> Mutex is an object that synchronizes access to a resource. It works as a locking mechanism that makes sure that only one thread can access the Mutex at a time. A semaphore is a signaling mechanism that uses atomic operations, wait and signal for synchronization. Semaphores allows multiple threads to access shared resources. There are two types of semaphores: counting (unrestricted value domain) and binary (values restricted to 0 and 1). 

### What are the differences between channels (in Communicating Sequential Processes, or as used by Go, Rust), mailboxes (in the Actor model, or as used by Erlang, D, Akka), and queues (as used by Python)? 
> A channel is symmetric and synchronous done via message passing, while a mailbox is asymmetric and asynchronous. A queue is FIFO. 

### List some advantages of using message passing over lock-based synchronization primitives.
> - Can store data in multiple places, meaning that message passing systems are suitable for larger systems with multiple data handlers. 
> - It is simpler to predict the worst-case execution time og the software with lock-free synchronization. 

### List some advantages of using lock-based synchronization primitives over message passing.
> - For simple, single systems, locking is more viable than messaging. 