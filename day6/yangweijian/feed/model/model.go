package model

import "camp/feed/model/feed"

func NewFeed() *feed.Feed {
	return new(feed.Feed)
}