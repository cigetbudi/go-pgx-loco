package util

import (
	"go-pgx-loco/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RetSucc(ctx *gin.Context, desc string, data interface{}) {
	r := model.Response{}
	r.StatusCode = "01"
	r.Description = desc
	r.Data = data
	ctx.JSON(http.StatusOK, r)
}

func RetBadReq(ctx *gin.Context, err error) {
	r := model.Response{}
	r.StatusCode = "01"
	r.Description = err.Error()
	ctx.JSON(http.StatusBadRequest, r)
}

func Timer(start time.Time, name string) {
	elapsed := time.Since(start)

	log.Printf("%s took %s", name, elapsed)
}
