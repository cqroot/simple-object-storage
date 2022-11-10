package common

import (
	"fmt"
	"path/filepath"
)

var storageRoot string = "/tmp/garden"

func GetBucketPath(account string, bucket string) string {
	return filepath.Join(storageRoot, account, fmt.Sprintf("%s.db", bucket))
}

func GetObjectPath(account string, bucket string, object string) string {
	return filepath.Join(
		storageRoot,
		account,
		bucket,
		object,
	)
}
