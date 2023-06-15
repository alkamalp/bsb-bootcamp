package api

import "github.com/gin-gonic/gin"

type Router interface {
	Route(handler *gin.RouterGroup)
}

type Api struct {
	server  *gin.Engine
	routers []Router
}

func Default() *Api {
	server := gin.Default()
	var routers = []Router{
		// parameter berupa struct Router pada module yang ingin di register ke main router
	}
	return &Api{server: server,
		routers: routers,
	}
}

func (a Api) Start() error {
	root := a.server.Group("/main/")
	for _, router := range a.routers {
		router.Route(root)
	}
	err := a.server.Run("0.0.0.0:3000")
	if err != nil {
		return err
	}
	return err
}
