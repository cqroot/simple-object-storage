package config

import "fmt"

func GetBucketServerAddress(account string, bucket string) (string, error) {
	return fmt.Sprintf("http://127.0.0.1:6002/v1/%s/%s", account, bucket), nil
}

func GetObjectServerAddress(account string, bucket string, object string) (string, error) {
	return fmt.Sprintf("http://127.0.0.1:6002/v1/%s/%s/%s", account, bucket, object), nil
}
