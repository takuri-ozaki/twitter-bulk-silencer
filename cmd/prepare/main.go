package main

import (
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"os"
	"twitter-bulk-silencer/internal/user"
	"twitter-bulk-silencer/internal/util"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("specify target(block/mute/follower/followee)")
		return
	}

	target := util.NewTarget(flag.Args()[0])
	if target == "unknown" {
		fmt.Println("specify target(block/mute/follower/followee)")
		return
	}

	err := godotenv.Load(fmt.Sprintf("./.env"))
	if err != nil {
		fmt.Println("prepare .env file")
		return
	}

	api := anaconda.NewTwitterApiWithCredentials(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))

	quiet := user.NewHandler(api, target, util.NewRealFileSystem(os.Getenv("BASE_DIR")))
	err = quiet.SaveUserList()
	if err != nil {
		panic(err)
	}
}
