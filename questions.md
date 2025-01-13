Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Concurrency is when multiple tasks are being worked on at the same time, but not necessarily running simultaneously. It’s more about managing several tasks by switching between them efficiently. Parallelism, on the other hand, is when multiple tasks are executed at the exact same time on different CPU cores. So, concurrency is about handling multiple tasks, while parallelism is about actually doing multiple things at once

What is the difference between a *race condition* and a *data race*? 
> A race condition happens when the program’s behavior depends on the unpredictable timing or order of operations, which can lead to incorrect behavior. A data race is a specific type of race condition where two or more threads access the same memory location at the same time, and at least one of them is writing to it without proper synchronization. In short, all data races are race conditions, but not all race conditions are data races. 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler manages how tasks are run on the CPU. It decides which threads or processes get to run and for how long. It balances the workload by switching between tasks, ensuring that high-priority tasks get more CPU time and that system resources are used efficiently. It does this through different strategies like round-robin scheduling or prioritizing tasks based on importance.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> Multiple threads are useful when a program needs to handle several tasks at once. They help improve performance by allowing different parts of a program to run in parallel, especially on multi-core CPUs. Threads are great for solving problems like handling user input while processing data, running background tasks, or managing network requests without freezing the program.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> Fibers and coroutines are lightweight alternatives to threads. They are managed by the program itself rather than the operating system. Coroutines, for example, can pause and resume their execution, which makes them more efficient for handling many small tasks, especially I/O operations. They use less memory and have faster context switching compared to threads, making them a better choice when scaling applications with thousands of tasks.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Honestly, both. It makes programs more efficient and responsive, but it also introduces more complexity. Handling things like race conditions, deadlocks, and synchronization can be really tricky and harder to debug. So, while concurrency can make programs better, it also requires careful design and testing.

What do you think is best - *shared variables* or *message passing*?
> I think it depends on the situation, but I personally prefer message passing because it’s safer and easier to reason about. Shared variables can be faster but come with a high risk of race conditions and require careful synchronization. Message passing, on the other hand, avoids those issues by keeping data isolated and communicating through channels or messages, which feels cleaner and more scalable to me.


