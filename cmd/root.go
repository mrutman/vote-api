package cmd

import (
	"fmt"

	"github.com/juju/loggo"

	"github.com/mrutman/vote-api/api"
	"github.com/mrutman/vote-api/api/v1alfa"
)

var logger = loggo.GetLogger("cmd")

func Run() {
	logger.SetLogLevel(loggo.INFO)
	fmt.Printf("start\n")

	logger.Errorf("run")
	logger.Infof("run")

	voteAPI := v1alfa.NewVoteAPI()

	voteServer := api.NewServer(voteAPI)
	if err := voteServer.RegisterAndServe(); err != nil {
		panic(err)
	}
}
