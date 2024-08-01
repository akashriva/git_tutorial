package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}
type routes struct{
	router *gin.Engine
}

type Routes []Route

func (r routes) EcommerceHealthCheck(rg *gin.RouterGroup) {
	orderRouteGrouping := rg.Group("/ecommerce")
	orderRouteGrouping.Use(CORSMiddleware())
	for _, route := range healthCheckRoutes {
		switch route.Method {
		case "GET":
			orderRouteGrouping.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			orderRouteGrouping.POST(route.Pattern, route.HandlerFunc)
		case "OPTIONS":
			orderRouteGrouping.OPTIONS(route.Pattern, route.HandlerFunc)
		case "PUT":
			orderRouteGrouping.OPTIONS(route.Pattern, route.HandlerFunc)
		case "DELETE":
			orderRouteGrouping.DELETE(route.Pattern, route.HandlerFunc)
		default:
			orderRouteGrouping.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid http method with this route",
				})
			})
		}

	}
}
func ClientRoutes() {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group(os.Getenv("API_VERSION"))
	r.EcommerceHealthCheck(v1)

	if err := r.router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Println("Failed to run server: ", err)
	}
}

//Middlewares

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Context-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == "OPTIONS" {
			ctx.Status(http.StatusOK)
		}
	}
}
