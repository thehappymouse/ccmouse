package engine

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestIsDuplicate(t *testing.T) {
	v1 := IsDuplicate("http://baidu.com/1.html")
	v2 := IsDuplicate("http://baidu.com/1.html")
	assert.Equal(t, v1, false)
	assert.Equal(t, v2, true)
}
