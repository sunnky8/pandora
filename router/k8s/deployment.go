package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/ielepro/pandora/k8s"
	"github.com/ielepro/pandora/render"
)

func DeploymentList(c *gin.Context)  {
	dp := &k8s.Deployment{}
	restList, err := dp.List()
	if err != nil{
		render.AppError(c, err.Error())
	}
	render.JSON(c, restList)
}

func Resources(c *gin.Context)  {
	metrics := &k8s.Metrics{}
	resources, err := metrics.NodeResourceList()
	if err != nil{
		render.AppError(c, err.Error())
	}
	render.JSON(c, resources)
}

