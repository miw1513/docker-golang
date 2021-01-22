package beer

import "github.com/gin-gonic/gin"

func NewRoutes(r *gin.Engine) {
	beerRoute := r.Group("/beer")

	beerRoute.GET("/", getAllBeer)

}

type Allber struct {
	Message string `json:"message"`
}

// getAllBeer  xxxx
// @Summary Retrieves users from mongodb
// @Description Get Users
// @Produce json
// @Success 200 {object} Allber
// @Router /beer [get]
func getAllBeer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get All Beer",
	})
}
