package common_test

import (
	"testing"

	"github.com/cqroot/simple-object-storage/internal/common"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	common.InitConfig("gateway-server")
	assert.Equal(t, 6000, viper.GetInt("gateway-server.port"))
}
