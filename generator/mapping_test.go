package generator

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMapping(t *testing.T) {

	assert.Equal(t, mapU64ToI64(mapI64ToU64(math.MinInt64)), int64(math.MinInt64))
	assert.Equal(t, mapU64ToI64(mapI64ToU64(-1000)), int64(-1000))
	assert.Equal(t, mapU64ToI64(mapI64ToU64(0)), int64(0))
	assert.Equal(t, mapU64ToI64(mapI64ToU64(1000)), int64(1000))
	assert.Equal(t, mapU64ToI64(mapI64ToU64(math.MaxInt64)), int64(math.MaxInt64))

	assert.Equal(t, mapU32ToI32(mapI32ToU32(math.MinInt32)), int32(math.MinInt32))
	assert.Equal(t, mapU32ToI32(mapI32ToU32(-1000)), int32(-1000))
	assert.Equal(t, mapU32ToI32(mapI32ToU32(0)), int32(0))
	assert.Equal(t, mapU32ToI32(mapI32ToU32(1000)), int32(1000))
	assert.Equal(t, mapU32ToI32(mapI32ToU32(math.MaxInt32)), int32(math.MaxInt32))

	assert.Equal(t, mapI64ToU64(mapU64ToI64(math.MaxUint64)), uint64(math.MaxUint64))
	assert.Equal(t, mapI64ToU64(mapU64ToI64(1000)), uint64(1000))
	assert.Equal(t, mapI64ToU64(mapU64ToI64(0)), uint64(0))

	assert.Equal(t, mapI32ToU32(mapU32ToI32(math.MaxUint32)), uint32(math.MaxUint32))
	assert.Equal(t, mapI32ToU32(mapU32ToI32(1000)), uint32(1000))
	assert.Equal(t, mapI32ToU32(mapU32ToI32(0)), uint32(0))
}
