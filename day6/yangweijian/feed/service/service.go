package service

import "camp/feed/service/feed"

func NewFeed() *feed.Feed {
	return new(feed.Feed)
}