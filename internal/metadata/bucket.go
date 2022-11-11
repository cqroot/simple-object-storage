package metadata

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cqroot/simple-object-storage/internal/config"
)

func PutObjectToBucket(account string, bucket string, object string) error {
	bucketAddress, err := config.GetBucketServerAddress(account, bucket)
	if err != nil {
		return err
	}
	targetUrl := fmt.Sprintf("%s/%s", bucketAddress, object)

	req, err := http.NewRequest("PUT", targetUrl, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}

func DeleteObjectFromBucket(account string, bucket string, object string) error {
	bucketAddress, err := config.GetBucketServerAddress(account, bucket)
	if err != nil {
		return err
	}
	targetUrl := fmt.Sprintf("%s/%s", bucketAddress, object)

	req, err := http.NewRequest("DELETE", targetUrl, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}
