package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/garden/common"
	"github.com/cqroot/garden/models"
)

func ListObjects(c *gin.Context) {
	objects, err := models.ListObjects(common.GetBucketPath(c.Param("account"), c.Param("bucket")))

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var result string
	for _, obj := range objects {
		result = fmt.Sprintf("%s%s\n", result, obj.Name)
	}

	c.String(http.StatusOK, "%s", result)
}

func PutObject(c *gin.Context) {
	err := models.PutObject(
		common.GetBucketPath(c.Param("account"), c.Param("bucket")),
		c.Param("object"),
	)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(200, gin.H{
			"account": c.Param("account"),
			"bucket":  c.Param("bucket"),
			"object":  c.Param("object"),
		})
	}
}

func DeleteObject(c *gin.Context) {
	err := models.DeleteObject(
		common.GetBucketPath(c.Param("account"), c.Param("bucket")),
		c.Param("object"),
	)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(200, gin.H{
			"account": c.Param("account"),
			"bucket":  c.Param("bucket"),
			"object":  c.Param("object"),
		})
	}
}
