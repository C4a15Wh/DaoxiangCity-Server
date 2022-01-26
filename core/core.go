package core

import (
	"dxcserver/common"
	"dxcserver/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"
)

var Config = common.LoadConfig()

const (
	Success            = 1
	InternalError      = -1
	PermissionError    = -2
	AuthError          = -3
	RequestParamsError = -4
	SkinServerError    = -5
)

func JSONFail(ctx *gin.Context, retCode int, errorMsg string) {

	ctx.JSON(http.StatusOK, gin.H{
		"Code":      retCode,
		"Status":    "Fail",
		"RequestID": uuid.NewV4(),
		"ErrorMsg":  errorMsg,
	})
}

func JSONSuccess(ctx *gin.Context, result interface{}) {

	ctx.JSON(http.StatusOK, gin.H{
		"Code":      Success,
		"Status":    "Success",
		"RequestID": uuid.NewV4(),
		"Result":    result,
	})
}

func ShowGitAddress(ctx *gin.Context) {
	JSONSuccess(ctx, Config.GitAddr)
}

func Update(ctx *gin.Context) {
	var Request model.UpdateParams
	err := ctx.ShouldBindBodyWith(&Request, binding.JSON)
	if err != nil {
		JSONFail(ctx, RequestParamsError, err.Error())
		log.Println("RequestParamsError: ", err)
		return
	}

	RequestHeader := []string{"content-type", "application/json", "Authorization", "Bearer" + Request.Token}

	var ServerResp model.SkinServerResp
	Resp, err := common.PostJson("https://daoxiangcity.com/api/auth/refresh", []byte(Request.Token), RequestHeader) // 中间这个参数可有可无

	if err != nil {
		JSONFail(ctx, SkinServerError, err.Error())
		log.Println("SkinServerError: ", err)
		return
	}

	err = json.Unmarshal(Resp, &ServerResp)

	if err != nil {
		JSONFail(ctx, InternalError, err.Error())
		log.Println("InternalError: ", err)
		return
	}

}

func Download(ctx *gin.Context) {

}
