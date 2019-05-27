package engine

import (
	"fmt"
	"testing"
)

func TestIsDuplicate(t *testing.T) {
	v1 := IsDuplicate("http://baidu.com/1.html")
	v2 := IsDuplicate("http://baidu.com/1.html")
	fmt.Println(v1, v2)

}
