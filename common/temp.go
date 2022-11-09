package common

import "path/filepath"

func GetLocalPath(account string, bucket string, object string) string {
	var storageRoot string = "/tmp/garden"
	return filepath.Join(
		storageRoot,
		account,
		bucket,
		object,
	)
}
