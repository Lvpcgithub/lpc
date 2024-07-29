package ants

import (
	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	r.GET("/task1", task1Handler)
	r.GET("/task2", task2Handler)
	r.GET("/task3", task3Handler)
}

func initRoutesWF(r *gin.Engine) {
	r.GET("/task4", task4Handler)
	r.GET("/task5", task5Handler)
	r.GET("/task6", task6Handler)
	r.GET("/statsWF", statsHandler)
}

func initRoutesNonBlock(r *gin.Engine) {
	r.GET("/task7", task7Handler)
	r.GET("/task8", task8Handler)
	r.GET("/statsNB", statsHandlerNonBlock)
}
func initRoutesPreAlloc(r *gin.Engine) {
	r.GET("/task9", task9Handler)
	r.GET("/task10", task10Handler)
	r.GET("/statsPA", statsHandlerPreAlloc)
}

/*func initRoutesPriority(r *gin.Engine) {
	r.GET("/task11", taskHandler1)
	r.GET("/task12", taskHandler2)
	r.GET("/task13", taskHandler3)
	r.GET("/statsP", statsHandlerP)
}*/
