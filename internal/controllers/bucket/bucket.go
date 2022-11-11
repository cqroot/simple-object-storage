package bucket

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/simple-object-storage/internal/common"
	"github.com/cqroot/simple-object-storage/internal/models"
)

func ListObjects(c *gin.Context) {
	var account string = c.Param("account")
	var bucket string = c.Param("bucket")

	objects, err := models.ListObjects(common.GetBucketPath(account, bucket))

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
	var account string = c.Param("account")
	var bucket string = c.Param("bucket")
	var object string = c.Param("object")[1:]

	err := models.PutObject(
		common.GetBucketPath(account, bucket), object,
	)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"account": account,
			"bucket":  bucket,
			"object":  object,
		})
	}
}

func DeleteObject(c *gin.Context) {
	var account string = c.Param("account")
	var bucket string = c.Param("bucket")
	var object string = c.Param("object")[1:]

	err := models.DeleteObject(
		common.GetBucketPath(account, bucket), object,
	)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"account": account,
			"bucket":  bucket,
			"object":  object,
		})
	}
}
