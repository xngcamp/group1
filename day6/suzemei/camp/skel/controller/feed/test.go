package feed

import (
	"encoding/json"
	"fmt"

	"camp/lib"
	"camp/skel/api"
	"camp/skel/test"
)

// Add 封装controller.Add操作
func Add(txt string) (err error) {
	c := &FeedC{}
	p := &AddReq{
		Txt: txt,
	}
	body, err := lib.TestPost(c.Add, p)
	if err != nil {
		return
	}

	resp := &struct {
		lib.Resp
		Data AddResp `json:"data"`
	}{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}
	if resp.Ret != lib.CodeOk {
		err = fmt.Errorf("resp ret invalid: %v, body: %s", resp, body)
		return
	}

	test.Sleep()
	return
}

// Del 封装controller.Del操作
func Del(id int,txt string) (err error) {
	c := &FeedC{}
	p := &DelReq{
		ID: id,
		Txt:txt,
	}
	body, err := lib.TestPost(c.Del, p)
	if err != nil {
		return
	}

	resp := &struct {
		lib.Resp
		Data AddResp `json:"data"`
	}{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}
	if resp.Ret != lib.CodeOk {
		err = fmt.Errorf("resp ret invalid: %v, body: %s", resp, body)
		return
	}

	test.Sleep()
	return
}

// Get 封装controller.Get操作
func Get(id int) (feedApi *api.Feed, err error) {
	c := &FeedC{}
	p := &GetReq{
		ID: id,
	}
	body, err := lib.TestPost(c.Get, p)
	if err != nil {
		return
	}

	resp := &struct {
		lib.Resp
		Data GetResp `json:"data"`
	}{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}
	if resp.Ret != lib.CodeOk {
		err = fmt.Errorf("resp ret invalid: %v, body: %s", resp, body)
		return
	}

	feedApi = resp.Data.Feed
	return
}

// GetSkel 返回一个全新的skel对象
func GetFeed() (feedApi *api.Feed) {
	id := test.GetID()
	for idTemp := id; ; idTemp++ {
		if skelTemp, err := Get(idTemp); err != nil {
			panic(err)
		} else if skelTemp != nil {
			if err := Del(idTemp); err != nil {
				panic(err)
			}
		} else {
			break
		}
	}

	if err := Add(id); err != nil {
		panic(err)
	}

	skelApi, err := Get(id)
	if err != nil {
		panic(err)
	} else if skelApi == nil {
		panic("get skel empty")
	}

	return
}

