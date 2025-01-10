# Results and Explanation

## Observed Behavior
When running the program, the result of the shared variable `i` is not always zero. Sometimes the result is a random number, either positive or negative.

## Why Does This Happen?

The variable `i` is accessed and modified by two threads concurrently:
- **Thread 1:** Increments `i` 1,000,000 times.  
- **Thread 2:** Decrements `i` 1,000,000 times.

The operations `i++` and `i--` are **not atomic**. Each of these operations involves three steps:
1. **Read** the current value of `i`.  
2. **Modify** the value (increment or decrement).  
3. **Write** the new value back to memory.  

Since both threads access and modify `i` simultaneously without synchronization, they can interfere with each other, causing some increments or decrements to be **lost**. This phenomenon is called a **race condition**.

## Expected Behavior

Without proper synchronization, the final value of `i` should not consistently be zero because the two threads are racing to update the same variable.

## How to Fix It

To ensure correct results, a synchronization mechanism such as a **mutex** can be used to lock the variable while one thread is updating it. This prevents both threads from modifying `i` at the same time.

Example fix using a `pthread_mutex_t`:
```c
pthread_mutex_lock(&lock);
i++;
pthread_mutex_unlock(&lock);
