package hello

import (
	"github.com/emicklei/go-restful"
)

// Resource is a resource for meta information about kublr.
type Resource struct {
}

// NewHelloResource creates new instance.
func NewResource() *Resource {
	return &Resource{}
}

// Register registers resource in restful container.
func (c *Resource) Register(container *restful.Container) *Resource {
	ws := new(restful.WebService)

	// MediaTypeApplicationYaml is a Mime Type for YAML
	const mediaTypeApplicationYaml = "application/x-yaml"

	ws.Path("/api/v1alfa"+"/hello").
		Doc("Hello api!!!").
		Consumes(restful.MIME_JSON, mediaTypeApplicationYaml).
		Produces(restful.MIME_JSON, mediaTypeApplicationYaml)

	ws.Route(ws.GET("").To(c.Hello).
		Doc("doc returns hello").
		Operation("operation returns hello").
		Param(ws.QueryParameter("cache", "set to 'nocache' to avoid cached data").DataType("string")))

	ws.Route(ws.GET("world").To(c.HelloWorld).
		Doc("doc returns hello world").
		Operation("operation returns hello world"))

	container.Add(ws)

	return c
}

// Hello returns hello
func (c *Resource) Hello(request *restful.Request, response *restful.Response) {
	response.WriteEntity("hello")
}

// HelloWprld returns hello world
func (c *Resource) HelloWorld(request *restful.Request, response *restful.Response) {
	response.WriteEntity("hello world")
}
