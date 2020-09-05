package util

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

type Api interface {
	GetSearch(string, url.Values) (anaconda.SearchResponse, error)
	BlockUserId(int64, url.Values) (anaconda.User, error)
	MuteUserId(int64, url.Values) (anaconda.User, error)
	GetBlocksIds(url.Values) (anaconda.Cursor, error)
	GetMutedUsersIds(url.Values) (anaconda.Cursor, error)
	GetFollowersIds(url.Values) (anaconda.Cursor, error)
	GetFriendsIds(url.Values) (anaconda.Cursor, error)
}

func NewRealApi(accessToken string, accessTokenSecret string, consumerKey string, consumerSecret string) Api {
	return anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)
}

type DummyApi struct {
}

func NewDummyAPI() Api {
	return &DummyApi{}
}

func (d *DummyApi) GetSearch(_ string, _ url.Values) (anaconda.SearchResponse, error) {
	return anaconda.SearchResponse{Statuses: []anaconda.Tweet{
		{User: anaconda.User{Id: 1001}, Text: "tweet1"},
		{User: anaconda.User{Id: 1002}, Text: "tweet2"},
		{User: anaconda.User{Id: 1003}, Text: "tweet3"},
		{User: anaconda.User{Id: 1004}, Text: "tweet4"},
		{User: anaconda.User{Id: 1005}, Text: "tweet5"},
		{User: anaconda.User{Id: 1001}, Text: "tweet6"},
		{User: anaconda.User{Id: 1001}, Text: "tweet7"},
		{User: anaconda.User{Id: 1002}, Text: "tweet8"},
		{User: anaconda.User{Id: 1003}, Text: "tweet9"},
		{User: anaconda.User{Id: 1006}, Text: "tweet10"},
		{User: anaconda.User{Id: 3}, Text: "block"},
		{User: anaconda.User{Id: 13}, Text: "mute"},
		{User: anaconda.User{Id: 23}, Text: "follower"},
		{User: anaconda.User{Id: 33}, Text: "followee"},
	}}, nil
}

func (d *DummyApi) BlockUserId(id int64, _ url.Values) (anaconda.User, error) {
	return anaconda.User{Id: id, IdStr: fmt.Sprint(id)}, nil
}

func (d *DummyApi) MuteUserId(id int64, _ url.Values) (anaconda.User, error) {
	return anaconda.User{Id: id, IdStr: fmt.Sprint(id)}, nil
}

func (d *DummyApi) GetBlocksIds(url.Values) (anaconda.Cursor, error) {
	return anaconda.Cursor{Ids: []int64{1, 2, 3, 4, 5}, Next_cursor: 0, Next_cursor_str: "0", Previous_cursor: 0, Previous_cursor_str: "0"}, nil
}

func (d *DummyApi) GetMutedUsersIds(url.Values) (anaconda.Cursor, error) {
	return anaconda.Cursor{Ids: []int64{11, 12, 13, 14, 15}, Next_cursor: 0, Next_cursor_str: "0", Previous_cursor: 0, Previous_cursor_str: "0"}, nil
}

func (d *DummyApi) GetFollowersIds(url.Values) (anaconda.Cursor, error) {
	return anaconda.Cursor{Ids: []int64{21, 22, 23, 24, 25}, Next_cursor: 0, Next_cursor_str: "0", Previous_cursor: 0, Previous_cursor_str: "0"}, nil
}

func (d *DummyApi) GetFriendsIds(url.Values) (anaconda.Cursor, error) {
	return anaconda.Cursor{Ids: []int64{31, 32, 33, 34, 35}, Next_cursor: 0, Next_cursor_str: "0", Previous_cursor: 0, Previous_cursor_str: "0"}, nil
}
