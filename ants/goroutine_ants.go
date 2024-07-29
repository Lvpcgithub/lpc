package ants

import (
	"github.com/gin-gonic/gin"
	"github.com/panjf2000/ants/v2"
	"log"
	"time"
)

var pool *ants.Pool
var poolP *ants.Pool
var poolNonBlock *ants.Pool
var poolPreAlloc *ants.Pool
var poolwf *ants.PoolWithFunc
var taskQueue *PriorityQueue

// GoroutineAnts 基础功能的pool
func GoroutineAnts() {
	var err error
	// 初始化10大小的协程池
	size := 10
	pool, err = ants.NewPool(size)
	if err != nil {
		log.Fatalf("failed to initialize pool: %v", err)
	}
	// 关闭协程池
	defer pool.Release()
	// 初始化route
	r := gin.Default()
	initRoutes(r)
	err1 := r.Run(":8080")
	if err1 != nil {
		return
	}
}

// GoroutineAntsDuration 超时时间的pool
func GoroutineAntsDuration() {
	var err error
	// 超时时间10s
	poolwf, err = ants.NewPoolWithFunc(10, func(payload interface{}) {
		handleTask(payload)
	}, ants.WithExpiryDuration(10*time.Second))
	if err != nil {
		log.Fatalf("failed to initialize pool: %v", err)
	}
	defer poolwf.Release()

	r := gin.Default()
	// Initialize routes
	initRoutesWF(r)

	err2 := r.Run(":8080")
	if err2 != nil {
		return
	}
}

func handleTask(payload interface{}) {
	if task, ok := payload.(func()); ok {
		task()
	} else {
		log.Println("received invalid task payload")
	}
}

// GoroutineAntsWithNonblocking 非阻塞功能的pool
func GoroutineAntsWithNonblocking() {
	var err error
	poolNonBlock, err = ants.NewPool(10, ants.WithNonblocking(true))
	if err != nil {
		log.Fatalf("failed to initialize pool: %v", err)
	}
	defer poolNonBlock.Release()
	r := gin.Default()
	// Initialize routes
	initRoutesNonBlock(r)
	r.Run(":8080")
}

// GoroutineAntsWithPreAlloc 预分配的pool
func GoroutineAntsWithPreAlloc() {
	var err error
	poolPreAlloc, err = ants.NewPool(10, ants.WithPreAlloc(true))
	if err != nil {
		log.Fatalf("failed to initialize pool: %v", err)
	}
	defer poolPreAlloc.Release()

	r := gin.Default()
	initRoutesPreAlloc(r)

	r.Run(":8080")
}

/*func GoroutineAntsPriority() {
	var err error
	// Initialize ants pool with a capacity of 10
	poolP, err = ants.NewPool(10)
	if err != nil {
		log.Fatalf("failed to initialize pool: %v", err)
	}
	defer poolP.Release()

	// Initialize priority queue
	taskQueue = &PriorityQueue{}
	heap.Init(taskQueue)

	r := gin.Default()
	// Initialize routes
	initRoutesPriority(r)

	// Start a worker to process tasks from the priority queue
	go processTasks()

	r.Run(":8080")
}

func processTasks() {
	for {
		if taskQueue.Len() > 0 {
			task := heap.Pop(taskQueue).(*Task)
			_ = pool.Submit(task.Execute)
		}
		time.Sleep(100 * time.Millisecond) // Adjust the sleep duration as needed
	}
}*/
