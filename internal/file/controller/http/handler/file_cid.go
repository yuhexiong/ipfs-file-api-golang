package handler

import (
	"bytes"
	"io"
	"ipfs-file-api/internal/file/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileCIDHandler struct {
	fileCIDService entity.FileCIDService
}

func NewFileCIDHandler(fileCIDService entity.FileCIDService) *FileCIDHandler {
	return &FileCIDHandler{
		fileCIDService: fileCIDService,
	}
}

func (h *FileCIDHandler) GetFileCID(ctx *gin.Context) {
	var dataUri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&dataUri); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := h.fileCIDService.GetFileCID(ctx, dataUri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *FileCIDHandler) CreateFileCID(ctx *gin.Context) {
	var dataUri struct {
		Name string `uri:"name" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&dataUri); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	photo, _, err := ctx.Request.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer photo.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, photo); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := h.fileCIDService.CreateFileCID(ctx, buf, dataUri.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
