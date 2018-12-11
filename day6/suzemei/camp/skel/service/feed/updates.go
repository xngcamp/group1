package feed

import (
	"camp/lib"
	"camp/skel/api"
)

// UpdateIndex 定义订阅者响应方法
func (feeds *Feeds) UpdateIndex(d interface{}, t lib.UpdateType, props map[string]interface{}) {
	switch o := d.(type) {
	case *api.Feed:
		switch t {
		case lib.ADD:
			_ = o
		case lib.DELETE:
			_ = o
		case lib.GET:
			_ = o
		}
	}
}

func init() {
	lib.RegisterIndex(&Feeds{})
}

