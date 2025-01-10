// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;

// Mutex for synchronization
pthread_mutex_t lock;
/*
Why Use a Mutex (pthread_mutex_t) Instead of a Semaphore?
A mutex is specifically designed to provide mutual exclusion for shared resources, making it the ideal choice for protecting the shared variable i.
Semaphores are more general and are better suited for signaling and managing access to multiple resources.
In this case, only one thread should modify i at a time, so a mutex is the most efficient and appropriate solution.
*/
// Note the return type: void*
void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    for (int j = 0; j < 1000000; j++) {
        pthread_mutex_lock(&lock);  // Acquire the lock
        i++;                        // Critical section
        pthread_mutex_unlock(&lock);  // Release the lock
    }
    return NULL;
}

void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times
    for (int j = 0; j<1000000; j++) {
        pthread_mutex_lock(&lock);  // Acquire the lock
        i--;                        // Critical section
        pthread_mutex_unlock(&lock);  // Release the lock
    }
    return NULL;
}


int main(){
    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?
    
    //Deklarer to variabler
    pthread_t thread1, thread2;

    // Initialize the mutex
    pthread_mutex_init(&lock, NULL);
    // Create threads
    pthread_create(&thread1, NULL, incrementingThreadFunction, &i);
    pthread_create(&thread2, NULL, decrementingThreadFunction, &i);
    
    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`   

 // pthread_join passer på at begge to treadsene er ferdige før man går videre. 
    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL); 

    
    printf("The magic number is: %d\n", i);
   
    // Destroy the mutex
    pthread_mutex_destroy(&lock);
    return 0;
}
