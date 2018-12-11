/*
Package service 用于定义服务层代码。
只允许在这里添加对外暴露的接口
*/
package service

import (
	"camp/skel/service/feed"
	"camp/skel/service/skel"
)

func NewSkel() *skel.Skel {
	return &skel.Skel{}
}
func NewFeed() *feed.Feeds {
	return &feed.Feeds{}
}
