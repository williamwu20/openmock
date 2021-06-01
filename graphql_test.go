package openmock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphql(t *testing.T) {
	om := &OpenMock{}
	om.SetRedis()

	t.Run("get non-exists data", func(t *testing.T) {
		v, err := redisHandleReply(om.redis.Do("get", "non-exist"))
		assert.NoError(t, err)
		assert.Empty(t, v)
	})
}
