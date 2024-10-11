Сетевой слой 
Этот слой никак не должен никогда связываться ни с какими стейтами, сторами и тд
Хотя в целом, может иметь некоторый свой стор для кеширования и тд 

Любой запрос приложения должен проходить через сетевой слой 
Запросы должны становиться в очередь 

Примеры:
1 - как не должно быть
```
useUpdateSomething {
    async mutation(data) {
        result = await fetch(..., body=data)
        if result ok {
            return
        } else {
            throw new Error "fetch_failed"
        }
        // etc
    }
    
    updateStore(data) {
        // save some data to store
        setSomeData(data)
    }

    async mutate(data) {
        await mutation(data)
        updateStore(data)
    }

    return {
        mutate
    }
}
```
Проблема этого кода в том, 
что прежде чем выполнится updateStore, 
будет задержка вызванная отправкой запроса в mutation.

2 - как это должно быть

```
class NetworkQueue {
    private queue: (() => Promise<void>)[] = [];
    private isProcessing = false;

    async processQueue() {
        if (this.isProcessing) return;
        this.isProcessing = true;

        while (this.queue.length > 0) {
            const task = this.queue.shift();
            if (task) {
                try {
                    await task(); // Process the task
                } catch (error) {
                    // Handle the error, e.g., retry logic
                    this.queue.push(task); // Re-add the task for retry
                }
            }
        }

        this.isProcessing = false;
    }

    addToQueue(task: () => Promise<void>) {
        this.queue.push(task);
        this.processQueue();
    }
}

const networkQueue = new NetworkQueue();

function useUpdateSomething() {
    async function mutation(data) {
        const result = await fetch('...', { method: 'POST', body: JSON.stringify(data) });
        if (!result.ok) {
            throw new Error('fetch_failed');
        }
    }
    
    async function updateStore(data) {
        // save some data to store asynchronously
        await setSomeData(data);
    }

    async function mutate(data) {
        // Update store immediately and add network task to the queue
        await Promise.any([
            updateStore(data),
            new Promise((resolve) => {
                networkQueue.addToQueue(async () => {
                    await mutation(data);
                    resolve(); // Resolve once the network task is successful
                });
            })
        ]);
    }

    return {
        mutate
    };
}
```

