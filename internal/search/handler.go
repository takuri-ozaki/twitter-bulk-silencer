package search

import (
	"net/url"
	"twitter-bulk-silencer/internal/util"
)

type Handler struct {
	api     util.Api
	keyword string
}

func NewHandler(api util.Api, keyword string) *Handler {
	return &Handler{api: api, keyword: keyword}
}

func (h *Handler) GetTargetUsers() ([]int64, error) {
	uv := url.Values{}
	uv.Set("count", "100")
	result, err := h.api.GetSearch(h.keyword, uv)
	if err != nil {
		return []int64{}, err
	}
	var list []int64
	for _, v := range result.Statuses {
		if !h.isContained(list, v.User.Id) {
			list = append(list, v.User.Id)
		}
	}
	return list, nil
}

func (h *Handler) isContained(userList []int64, user int64) bool {
	for _, v := range userList {
		if v == user {
			return true
		}
	}
	return false
}
