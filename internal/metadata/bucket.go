package metadata

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetBucketServerAddress(account string, bucket string) string {
	return fmt.Sprintf("http://127.0.0.1:6002/v1/%s/%s", account, bucket)
}

func PutObjectToBucket(account string, bucket string, object string) error {
	bucketAddress := GetBucketServerAddress(account, bucket)
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
	bucketAddress := GetBucketServerAddress(account, bucket)
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
