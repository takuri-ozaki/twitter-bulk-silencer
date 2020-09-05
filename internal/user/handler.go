package user

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"strconv"
	"twitter-bulk-silencer/internal/util"
)

type Handler struct {
	api        util.Api
	target     util.RealTarget
	FileSystem util.FileSystem
}

func NewHandler(api util.Api, target util.RealTarget, fileSystem util.FileSystem) *Handler {
	return &Handler{api: api, target: target, FileSystem: fileSystem}
}

func (h *Handler) SaveUserList() error {
	list, err := h.getUserList()
	if err != nil {
		return err
	}
	return h.writeUserList(list)
}

func (h *Handler) WriteAppendUserList(list []int64) error {
	file, err := h.FileSystem.OpenFile(h.target.GetFileName(), true)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, v := range list {
		fmt.Fprintln(file, v)
	}
	return nil
}

func (h *Handler) ReadUserList() ([]int64, error) {
	file, err := h.FileSystem.OpenFile(h.target.GetFileName(), false)
	if err != nil {
		return []int64{}, err
	}
	defer file.Close()

	var list []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		asInt, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			continue
		}
		list = append(list, asInt)
	}
	return list, nil
}

func (h *Handler) getUserList() ([]int64, error) {
	var list []int64
	var uv url.Values = nil
	if h.target == "follower" || h.target == "followee" {
		uv = url.Values{}
		uv.Set("count", "5000")
	}

	for {
		result, err := h.getUserListFunction()(uv)
		if err != nil {
			return list, err
		}
		for _, v := range result.Ids {
			list = append(list, v)
		}
		fmt.Println(strconv.Itoa(len(list)) + " user listed (" + h.target.String() + ")")

		if result.Next_cursor == 0 {
			break
		}

		uv = url.Values{}
		uv.Set("cursor", result.Next_cursor_str)
		if h.target == "follower" || h.target == "followee" {
			uv.Set("count", "5000")
		}
	}

	return list, nil
}

func (h *Handler) getUserListFunction() func(url.Values) (anaconda.Cursor, error) {
	switch h.target {
	case "block":
		return h.api.GetBlocksIds
	case "mute":
		return h.api.GetMutedUsersIds
	case "follower":
		return h.api.GetFollowersIds
	case "followee":
		return h.api.GetFriendsIds
	default:
		return h.dummyList
	}
}

func (h *Handler) writeUserList(list []int64) error {
	file, err := h.FileSystem.OpenFile(h.target.GetFileName(), false)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, v := range list {
		fmt.Fprintln(file, v)
	}
	return nil
}

func (h *Handler) dummyList(_ url.Values) (c anaconda.Cursor, err error) {
	return anaconda.Cursor{Previous_cursor_str: "0", Next_cursor_str: "0", Ids: []int64{}}, errors.New("dummy list called")
}
