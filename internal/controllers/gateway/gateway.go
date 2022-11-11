package gateway

import (
	"net/http"

	"github.com/cqroot/simple-object-storage/internal/config"
	"github.com/gin-gonic/gin"
)

func GetBucket(c *gin.Context) {
	var account string = c.Param("account")
	var bucket string = c.Param("bucket")

	targetUrl, err := config.GetBucketServerAddress(account, bucket)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req, err := http.NewRequest("GET", targetUrl, c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(resp.StatusCode, resp.ContentLength, "text/plain", resp.Body, nil)
}

func GetObject(c *gin.Context) {
	var account string = c.Param("account")
	var bucket string = c.Param("bucket")
	var object string = c.Param("object")[1:]

	targetUrl, err := config.GetObjectServerAddress(account, bucket, object)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req, err := http.NewRequest("GET", targetUrl, c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(resp.StatusCode, resp.ContentLength, "application/json", resp.Body, nil)
}

func PutObject(c *gin.Context) {
	var account string = c.Param("account")
	var bucket string = c.Param("bucket")
	var object string = c.Param("object")[1:]

	targetUrl, err := config.GetObjectServerAddress(account, bucket, object)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req, err := http.NewRequest("PUT", targetUrl, c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(resp.StatusCode, resp.ContentLength, "application/json", resp.Body, nil)
}

func DeleteObject(c *gin.Context) {
	var account string = c.Param("account")
	var bucket string = c.Param("bucket")
	var object string = c.Param("object")[1:]

	targetUrl, err := config.GetObjectServerAddress(account, bucket, object)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req, err := http.NewRequest("DELETE", targetUrl, c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(resp.StatusCode, resp.ContentLength, "application/json", resp.Body, nil)
}
