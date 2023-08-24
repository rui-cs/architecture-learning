package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rui-cs/architecture-learning/service"
)

type UserHandler struct {
	us *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{us: us}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/usermgt")

	ug.POST("/add", u.Add)
	ug.POST("/del", u.Del)
	ug.POST("/upd", u.Upd)

	ug.POST("/getByCompany", u.GetByCompany)
	ug.POST("/getByID", u.GetByID)
}

func (u *UserHandler) Add(ctx *gin.Context) {
	var req service.UserReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	err := u.us.Add(&req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "服务错误")
		return
	}

	ctx.String(http.StatusOK, "添加成功")
}

func (u *UserHandler) Del(ctx *gin.Context) {
	var req service.UserReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	err := u.us.Del(&req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "服务错误")
		return
	}

	ctx.String(http.StatusOK, "删除成功")
}

func (u *UserHandler) Upd(ctx *gin.Context) {
	var req service.UserReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	err := u.us.Upd(&req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "服务错误")
		return
	}

	ctx.String(http.StatusOK, "更新成功")
}

func (u *UserHandler) GetByCompany(ctx *gin.Context) {
	var req service.UserReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	get, err := u.us.Get(&req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "服务错误")
		return
	}

	ctx.JSON(http.StatusOK, get)
}

func (u *UserHandler) GetByID(ctx *gin.Context) {
	var req service.UserReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	get, err := u.us.Get(&req, req.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "服务错误")
		return
	}

	ctx.JSON(http.StatusOK, get)
}
