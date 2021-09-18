package app

import (
	"github.com/ayush723/to-do-list/src/controllers/list"
	"github.com/ayush723/to-do-list/src/controllers/ping"
	"github.com/gin-gonic/gin"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)               //done
	router.POST("/todo", list.Create)            //done
	router.GET("/todo", list.Get)                //done
	router.GET("/todo/:todo-id", list.GetById)   //done
	router.PATCH("/todo/:todo-id", list.Update)  // done
	router.DELETE("/todo/:todo-id", list.Delete) //done
	router.GET("/todo/search", list.Search)      // done

}
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
