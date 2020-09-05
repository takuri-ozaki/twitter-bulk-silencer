package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
	"twitter-bulk-silencer/internal/search"
	"twitter-bulk-silencer/internal/silence"
	"twitter-bulk-silencer/internal/util"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("specify search text")
		return
	}
	execute := false
	if flag.Arg(1) == "execute" {
		execute = true
	}

	err := godotenv.Load(fmt.Sprintf("./.env"))
	if err != nil {
		panic(err)
	}

	enableStandardQuery, err := strconv.ParseBool(os.Getenv("ENABLE_STANDARD_QUERY"))
	if err != nil {
		panic(err)
	}
	if !enableStandardQuery && !strings.HasPrefix(flag.Arg(0), "#") {
		fmt.Println("search text should be hashtag or enable standard query")
		return
	}

	blockMode, err := strconv.ParseBool(os.Getenv("BLOCK_MODE"))
	if err != nil {
		panic(err)
	}
	protectFollower, err := strconv.ParseBool(os.Getenv("PROTECT_FOLLOWER"))
	if err != nil {
		panic(err)
	}
	protectFollowee, err := strconv.ParseBool(os.Getenv("PROTECT_FOLLOWEE"))
	if err != nil {
		panic(err)
	}

	api := util.NewRealApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))

	searcher := search.NewHandler(api, flag.Arg(0))
	silencer := silence.NewHandler(api, blockMode, 5, protectFollower, protectFollowee, execute, util.NewRealFileSystem(os.Getenv("BASE_DIR")))

	targetList, err := searcher.GetTargetUsers()
	if err != nil {
		panic(err)
	}
	err = silencer.Silence(targetList)
	if err != nil {
		panic(err)
	}
}
