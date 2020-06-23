package v1alfa

import (
	"log"
	"os"

	//	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"

	"github.com/mrutman/vote-api/api/v1alfa/hello"
)

var logger = loggo.GetLogger("KublrAPI")

// VoteAPI is a definition of Vote API.
type VoteAPI struct {
}

// NewVoteAPI creates new instance of Vote API.
// It is required to call Register before start to use it.
func NewVoteAPI() *VoteAPI {

	api := &VoteAPI{}

	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.SetLogger(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds))

	return api
}

// Register registers REST resources in container.
func (api *VoteAPI) Register(wsContainer *restful.Container, insecure bool) error {
	hello.NewResource().Register(wsContainer)
	return nil
}
