package object

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/simple-object-storage/internal/common"
)

func PutObject(c *gin.Context) {
	storagePath := common.GetObjectPath(
		c.Param("account"),
		c.Param("bucket"),
		c.Param("object"),
	)

	err := os.MkdirAll(filepath.Dir(storagePath), os.ModePerm)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	f, err := os.Create(storagePath)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"account": c.Param("account"),
		"bucket":  c.Param("bucket"),
		"object":  c.Param("object"),
	})
}

func GetObject(c *gin.Context) {
	storagePath := common.GetObjectPath(
		c.Param("account"),
		c.Param("bucket"),
		c.Param("object"),
	)

	c.File(storagePath)
}

func DeleteObject(c *gin.Context) {
	storagePath := common.GetObjectPath(
		c.Param("account"),
		c.Param("bucket"),
		c.Param("object"),
	)

	err := os.Remove(storagePath)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"account": c.Param("account"),
				"bucket":  c.Param("bucket"),
				"object":  c.Param("object"),
			})
			return
		} else {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
	}

	c.JSON(200, gin.H{
		"account": c.Param("account"),
		"bucket":  c.Param("bucket"),
		"object":  c.Param("object"),
	})
}
