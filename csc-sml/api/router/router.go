package router

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		UserGroup(v1)
	}

}
