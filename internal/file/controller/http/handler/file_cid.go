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

// @Summary 取得指定 File
// @Description 取得指定 File
// @Tags FileCID
// @Produce application/octet-stream
// @Security OAuth2Password
// @Param id path integer true "FileCID ID"
// @Success 200 {string} binary "File"
// @Failure 400 ""
// @Failure 500 ""
// @Router /api/file-cid/{id} [get]
func (h *FileCIDHandler) GetFileCID(ctx *gin.Context) {
	var dataUri struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&dataUri); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	buf, err := h.fileCIDService.GetFileCID(ctx, dataUri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	contentType := "application/octet-stream"
	ctx.Data(http.StatusOK, contentType, *buf)
}

// @Summary 上傳 File
// @Description 上傳 File
// @Tags FileCID
// @Produce json
// @Security OAuth2Password
// @Param name path string true "File Name"
// @Param file formData file true "File"
// @Success 201 {object} entity.FileCID
// @Failure 400 ""
// @Failure 500 ""
// @Router /api/file-cid/{name} [post]
func (h *FileCIDHandler) CreateFileCID(ctx *gin.Context) {
	var dataUri struct {
		Name string `uri:"name" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&dataUri); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
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
