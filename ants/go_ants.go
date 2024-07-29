package ants

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func task1Handler(c *gin.Context) {
	err := pool.Submit(func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task 1 completed")
	})
	if err != nil {
		log.Printf("failed to submit task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task 1 is being processed"})
}

func task2Handler(c *gin.Context) {
	err := pool.Submit(func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Task 2 completed")
	})
	if err != nil {
		log.Printf("failed to submit task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task 2 is being processed"})
}

func task3Handler(c *gin.Context) {
	err := pool.Submit(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Task 3 completed")
	})
	if err != nil {
		log.Printf("failed to submit task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task 3 is being processed"})
}

func task4Handler(c *gin.Context) {
	err := poolwf.Invoke(func() {
		// Simulate a task that takes some time
		time.Sleep(2 * time.Second)
		fmt.Println("Task 4 completed")
	})
	if err != nil {
		log.Printf("failed to submit task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task 4 is being processed"})
}

func task5Handler(c *gin.Context) {
	err := poolwf.Invoke(func() {
		// Simulate a task that takes some time
		time.Sleep(22 * time.Second)
		fmt.Println("Task 5 completed")
	})
	if err != nil {
		log.Printf("failed to submit task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task 5 is being processed"})
}

func task6Handler(c *gin.Context) {
	err := poolwf.Invoke(func() {
		// Simulate a task that takes some time
		time.Sleep(1 * time.Second)
		fmt.Println("Task 6 completed")
	})
	if err != nil {
		log.Printf("failed to submit task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task 6 is being processed"})
}

func statsHandler(c *gin.Context) {
	running := poolwf.Running()
	capacity := poolwf.Cap()
	c.JSON(http.StatusOK, gin.H{
		"running":  running,
		"capacity": capacity,
	})
}

func task7Handler(c *gin.Context) {
	err := poolNonBlock.Submit(func() {
		// Simulate a task that takes some time
		time.Sleep(7 * time.Second)
		fmt.Println("Task 7 completed")
	})
	if err != nil {
		log.Printf("failed to submit task: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "too many requests"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task 7 is being processed"})
}

func task8Handler(c *gin.Context) {
	err := poolNonBlock.Submit(func() {
		// Simulate a task that takes some time
		time.Sleep(8 * time.Second)
		fmt.Println("Task 8 completed")
	})
	if err != nil {
		log.Printf("failed to submit task: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "too many requests"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task 8 is being processed"})
}

func statsHandlerNonBlock(c *gin.Context) {
	running := poolNonBlock.Running()
	capacity := poolNonBlock.Cap()
	free := capacity - running
	c.JSON(http.StatusOK, gin.H{
		"running":  running,
		"capacity": capacity,
		"free":     free,
	})
}

func task9Handler(c *gin.Context) {
	err := poolPreAlloc.Submit(func() {
		// Simulate a task that takes some time
		time.Sleep(7 * time.Second)
		fmt.Println("Task 9 completed")
	})
	if err != nil {
		log.Printf("failed to submit task: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "too many requests"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task 9 is being processed"})
}

func task10Handler(c *gin.Context) {
	err := poolPreAlloc.Submit(func() {
		// Simulate a task that takes some time
		time.Sleep(8 * time.Second)
		fmt.Println("Task 10 completed")
	})
	if err != nil {
		log.Printf("failed to submit task: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "too many requests"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task 10 is being processed"})
}

func statsHandlerPreAlloc(c *gin.Context) {
	running := poolPreAlloc.Running()
	capacity := poolPreAlloc.Cap()
	free := capacity - running
	c.JSON(http.StatusOK, gin.H{
		"running":  running,
		"capacity": capacity,
		"free":     free,
	})
}

/*func taskHandler1(c *gin.Context) {
	priority := 1
	task := &Task{
		Priority: priority,
		Execute: func() {
			time.Sleep(12 * time.Second)
			log.Println("Task 1 completed with priority", priority)
		},
	}
	heap.Push(taskQueue, task)
	c.JSON(http.StatusOK, gin.H{"message": "Task 1 is being processed"})
}

func taskHandler2(c *gin.Context) {
	priority := 2
	task := &Task{
		Priority: priority,
		Execute: func() {
			time.Sleep(13 * time.Second)
			log.Println("Task 2 completed with priority", priority)
		},
	}
	heap.Push(taskQueue, task)
	c.JSON(http.StatusOK, gin.H{"message": "Task 2 is being processed"})
}

func taskHandler3(c *gin.Context) {
	priority := 3
	task := &Task{
		Priority: priority,
		Execute: func() {
			time.Sleep(14 * time.Second)
			log.Println("Task 3 completed with priority", priority)
		},
	}
	heap.Push(taskQueue, task)
	c.JSON(http.StatusOK, gin.H{"message": "Task 3 is being processed"})
}

func statsHandlerP(c *gin.Context) {
	running := pool.Running()
	capacity := pool.Cap()
	free := capacity - running
	c.JSON(http.StatusOK, gin.H{
		"running":  running,
		"capacity": capacity,
		"free":     free,
	})
	log.Printf("Pool Stats - Running: %d, Capacity: %d, Free: %d", running, capacity, free)
}
*/
