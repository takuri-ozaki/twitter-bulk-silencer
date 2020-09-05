package silence

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"strconv"
	"twitter-bulk-silencer/internal/user"
	"twitter-bulk-silencer/internal/util"
)

type Handler struct {
	api             util.Api
	mode            util.Mode
	errorLimit      int
	protectFollower bool
	protectFollowee bool
	execute         bool
	silenceeHandler *user.Handler
	followerHandler *user.Handler
	followeeHandler *user.Handler
}

func NewHandler(api util.Api, blockMode bool, errorLimit int, protectFollower bool, protectFollowee bool, execute bool, fileSystem util.FileSystem) *Handler {
	mode := "mute"
	if blockMode {
		mode = "block"
	}
	return &Handler{
		api:             api,
		mode:            util.NewMode(mode),
		errorLimit:      errorLimit,
		protectFollower: protectFollower,
		protectFollowee: protectFollowee,
		execute:         execute,
		silenceeHandler: user.NewHandler(api, util.NewTarget(mode), fileSystem),
		followerHandler: user.NewHandler(api, util.NewTarget("follower"), fileSystem),
		followeeHandler: user.NewHandler(api, util.NewTarget("followee"), fileSystem),
	}
}

func (h *Handler) Silence(targetList []int64) error {
	var silencedList []anaconda.User
	targetList, err := h.filter(targetList)

	if err != nil {
		return err
	}
	errorCount := 0
	for _, v := range targetList {
		silenced, err := h.getSilenceFunction()(v, nil)
		if err != nil {
			errorCount++
			if errorCount > h.errorLimit {
				return err
			}
			continue
		}
		silencedList = append(silencedList, silenced)
		fmt.Println(silenced.ScreenName + " (" + silenced.IdStr + ") was silenced")
	}
	fmt.Println("----------")
	fmt.Println(strconv.Itoa(len(silencedList)) + " users are silenced")
	if h.execute {
		return h.silenceeHandler.WriteAppendUserList(h.extract(silencedList))
	} else {
		return nil
	}
}

func (h *Handler) getSilenceFunction() func(int64, url.Values) (anaconda.User, error) {
	if !h.execute {
		return h.dryRunUserIdFunction
	}
	if h.mode.IsBlockMode() {
		return h.api.BlockUserId
	}
	return h.api.MuteUserId
}

func (h *Handler) filter(list []int64) ([]int64, error) {
	var list1 []int64
	var err error
	if h.protectFollower {
		list1, err = h.getDiff(list, h.followerHandler, " was skipped (follower)")
		if err != nil {
			return list, err
		}
	} else {
		list1 = list
	}

	var list2 []int64
	if h.protectFollowee {
		list2, err = h.getDiff(list1, h.followeeHandler, " was skipped (followee)")
		if err != nil {
			return list, err
		}
	} else {
		list2 = list1
	}

	list3, err := h.getDiff(list2, h.silenceeHandler, "")
	if err != nil {
		return list, err
	}

	return list3, nil
}

func (h *Handler) getDiff(input []int64, handler *user.Handler, message string) ([]int64, error) {
	var output []int64
	userList, err := handler.ReadUserList()
	if err != nil {
		return input, err
	}
	for _, v1 := range input {
		isFound := false
		for _, v2 := range userList {
			if v1 == v2 {
				isFound = true
				if len(message) != 0 {
					fmt.Println(fmt.Sprint(v2) + message)
				}
				continue
			}
		}
		if !isFound {
			output = append(output, v1)
		}
	}
	return output, nil
}

func (h *Handler) extract(userList []anaconda.User) []int64 {
	var idList []int64
	for _, v := range userList {
		idList = append(idList, v.Id)
	}
	return idList
}

func (h *Handler) dryRunUserIdFunction(id int64, _ url.Values) (anaconda.User, error) {
	return anaconda.User{Id: id, IdStr: fmt.Sprint(id), ScreenName: "?????"}, nil
}
