#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t counter_lock;

//POSIX has both mutexes and semaphoes.
//In this case I would use a mutex, as it makes no sense that both threads will change the variable at the same time.
//Then we can not use a semaphore.

// Note the return type: void*
void* incrementingThreadFunction(){
    for (int j = 0; j < 1000000; j++) {
        // TODO: sync access to i
        pthread_mutex_lock(&counter_lock);
        i++;
        pthread_mutex_unlock(&counter_lock);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    for (int j = 0; j < 1000000; j++) {
        // TODO: sync access to i
        pthread_mutex_lock(&counter_lock);
        i--;
        pthread_mutex_unlock(&counter_lock);
    }
    return NULL;
}


int main(){
    pthread_mutex_init(&counter_lock,NULL);
    
    pthread_t incrementingThread, decrementingThread;
    
    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    
    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);
    
    pthread_mutex_destroy(&counter_lock);
    
    printf("The magic number is: %d\n", i);
    return 0;
}
