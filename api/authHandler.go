package api

import (
	"go-pgx-loco/model"
	"go-pgx-loco/repository"
	"go-pgx-loco/util"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	u := model.User{}

	if err := ctx.ShouldBindJSON(&u); err != nil {
		util.RetBadReq(ctx, err)
		return
	}

	err := repository.AddUser(&u)
	if err != nil {
		util.RetBadReq(ctx, err)
		return
	}
	util.RetSucc(ctx, "berhasil mendaftar akun", nil)
}
