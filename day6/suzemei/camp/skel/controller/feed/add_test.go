package feed


import (
	"testing"

	"camp/skel/test"
)

// 测试/skel/add
func TestAdd(t *testing.T) {
	test.Setup()

	skel := GetSkel()
	if skel == nil {
		t.Fatal("test fail")
	}
}
