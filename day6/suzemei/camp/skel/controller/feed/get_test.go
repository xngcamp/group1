package feed

import (
	"testing"

	"camp/skel/test"
)

// 测试/skel/get
func TestGet(t *testing.T) {
	test.Setup()

	feed:= GetFeed()
	if feed== nil {
		t.Fatal("test fail")
	}
}

