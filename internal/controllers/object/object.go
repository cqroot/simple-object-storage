package object

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/simple-object-storage/internal/common"
	"github.com/cqroot/simple-object-storage/internal/metadata"
)

func GetObject(c *gin.Context) {
	var account string = c.Param("account")
	var bucket string = c.Param("bucket")
	var object string = c.Param("object")[1:]

	storagePath := common.GetObjectPath(account, bucket, object)

	c.File(storagePath)
}

func PutObject(c *gin.Context) {
	var account string = c.Param("account")
	var bucket string = c.Param("bucket")
	var object string = c.Param("object")[1:]

	storagePath := common.GetObjectPath(account, bucket, object)

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

	err = metadata.PutObjectToBucket(account, bucket, object)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func DeleteObject(c *gin.Context) {
	var account string = c.Param("account")
	var bucket string = c.Param("bucket")
	var object string = c.Param("object")[1:]

	storagePath := common.GetObjectPath(account, bucket, object)

	err := os.Remove(storagePath)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"account": account,
				"bucket":  bucket,
				"object":  object,
			})
			return
		} else {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
	}

	err = metadata.DeleteObjectFromBucket(account, bucket, object)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
