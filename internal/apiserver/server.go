package apiserver

import (
	"ACAutomatic/internal/pkg/ac"
	"ACAutomatic/pkg/fs"
	"ACAutomatic/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

var automatic *ac.AC

func init() {
	data, err := fs.LoadConfig("config/config.txt")
	if err != nil {
		panic(err)
	}
	automatic = ac.NewAC(data)
	log.Info("init acAutomatic success!!!")
}

type Server struct {
	engine *gin.Engine
}

func CreateServer() (*gin.Engine, error) {
	engine := gin.Default()
	engine.Use(Cors())

	engine.POST("/replace", ContentType, Replace)
	engine.POST("/delete", Delete)
	engine.POST("/insert", Insert)

	return engine, nil
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}

func ContentType(ctx *gin.Context) {
	if ctx.ContentType() != gin.MIMEJSON {
		ctx.AbortWithStatusJSON(200, gin.H{"msg": "only support application/json"})
		ctx.Abort()
	}
	ctx.Next()
}

func Replace(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	log.Info(string(data))
	log.Info(automatic.Replace(string(data)))
	ctx.String(200, automatic.Replace(string(data)))
}

func Insert(ctx *gin.Context) {
	_line, _ := ctx.GetRawData()
	line := string(_line)
	automatic.Insert(line)
	fs.Insert(line, "config/config.txt")
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "插入成功",
	})
}

func Delete(ctx *gin.Context) {
	_line, _ := ctx.GetRawData()
	line := string(_line)
	isDelete := automatic.Delete(line)
	if isDelete {
		fs.Delete(line, "config/config.txt")
	} else {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除失败，删除的字符处于trie树需要复用。",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
