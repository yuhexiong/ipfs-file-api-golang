package route

import (
	"ipfs-file-api/internal/file/controller/http/handler"

	"github.com/gin-gonic/gin"
)

var _ Router = (*Route)(nil)

type Router interface {
	RouterGroup(g *gin.RouterGroup)
}

type Route struct {
	FileCIDHandler *handler.FileCIDHandler
}

func NewRouter(fileCIDHandler *handler.FileCIDHandler) Router {
	return &Route{FileCIDHandler: fileCIDHandler}
}

func (r *Route) RouterGroup(g *gin.RouterGroup) {
	g.GET("/file-cid/:id", r.FileCIDHandler.GetFileCID)
	g.POST("/file-cid/:name", r.FileCIDHandler.CreateFileCID)
}
