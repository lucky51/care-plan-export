package main

import (
	"bytes"
	"fmt"
	_ "goexcel/config"
	_ "goexcel/docs"
	_ "goexcel/log"
	"goexcel/modals"
	"goexcel/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	//xlst "github.com/ivahaev/go-xlsx-templater"
	xlst "gitee.com/lucky51/go-xlsx-templater"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"     //swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.uber.org/zap"
)

// GetExcelExportHandlerFunc 照护执行计划查询并导出excel
// @Summary 照护执行计划查询并导出excel
// @version 0.0.1
// @description  输入参数大小写敏感.
// @BasePath /
// @Produce  application/octet-stream
// @Param residentId query int true "住户ID",
// @Param tenantId query int true "租户ID",
// @Param startDate query string true "开始日期",
// @Param endDate query string true "结束日期",
// @Success 200 {string} json
// @Router /file [get]
func GetExcelExportHandlerFunc(context *gin.Context) {
	jsonInput := &services.CareResidentRequestInput{}
	bindErr := context.ShouldBindQuery(jsonInput) //BindQuery()
	if bindErr != nil {
		context.JSON(http.StatusInternalServerError, bindErr.Error())
		context.Abort()
		return
	}
	fmt.Printf("%+v \r\n", jsonInput)
	response, e := services.CareResidentRecordGroupByWeek(jsonInput)
	if e != nil {
		context.AbortWithError(http.StatusInternalServerError, e)
		return
	}
	if !response.Success {
		context.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "远程服务调用失败",
			"error":   response.Error,
		})
		return
	}
	if len(response.Result.Data) == 0 {
		context.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "未查到任何数据",
		})
		return
	}
	outerParams := modals.MapTo(&response.Result.Data)
	doc := xlst.New()
	doc.ReadTemplate("./excel/template-new.xlsx")
	renderErr := doc.RenderMultipleSheet(nil, outerParams...)
	if renderErr != nil {
		context.AbortWithError(http.StatusInternalServerError, renderErr)
		return
	}
	var b bytes.Buffer
	if wErr := doc.Write(&b); wErr != nil {
		context.AbortWithError(http.StatusInternalServerError, wErr)
		return
	}

	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	context.Header("Content-Description", "File Transfer")
	context.Header("Content-Disposition", "attachment; filename="+downloadName)
	context.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}

// PostJSONDataHandlerFunc 照护执行计划excel生成-服务端转发
// @Summary 照护执行计划excel生成-服务端转发
// @Version 0.0.1
// @BasePath /
// @Accept application/json
// @Produce application/json
// @Param postJson body []modals.OuterInput true "JSON"
// @Success 200 {string} json
// @Router /file [post]
func PostJSONDataHandlerFunc(context *gin.Context) {
	jsonInput := new([]modals.OuterInput)
	err := context.ShouldBindJSON(jsonInput)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if len(*jsonInput) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请传入有效数据",
		})
		return
	}

	outerParams := modals.MapTo(jsonInput)

	doc := xlst.New()

	doc.ReadTemplate("./excel/template-new.xlsx")
	renderErr := doc.RenderMultipleSheet(nil, outerParams...)

	if renderErr != nil {
		context.AbortWithError(http.StatusInternalServerError, renderErr)
		return
	}
	var b bytes.Buffer
	if wErr := doc.Write(&b); wErr != nil {
		context.AbortWithError(http.StatusInternalServerError, wErr)
		return
	}
	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	context.Header("Content-Description", "File Transfer")
	context.Header("Content-Disposition", "attachment; filename="+downloadName)
	context.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}

// IndexPageHandlerFunc 首页控制器
func IndexPageHandlerFunc(listenPort *int, env *string) func(ctx *gin.Context) {
	return func(context *gin.Context) {
		appName := viper.GetString("app")
		if appName == "" {
			appName = "hello gin"
		}
		e := services.PingCmp()
		cmpStatus := "closed"
		if e == nil {
			cmpStatus = "connected"
		}
		cmpURL := viper.GetString("cmpURL")
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":  appName,
			"port":   listenPort,
			"cmpURL": cmpURL,
			"env":    env,
			"pingAt": time.Now().Format("15:04:05"),
			"status": cmpStatus,
		})
	}
}

func main() {
	env := viper.GetString("env")
	if env != "" {
		gin.SetMode(env)
	}
	listenPort := viper.GetInt("port")
	if listenPort == 0 {
		listenPort = 3001
	}

	cmpURL := viper.GetString("cmpURL")
	zap.L().Info("viper setting:",
		zap.String("env", env),
		zap.Int("port", listenPort),
		zap.String("cmpURL", cmpURL),
	)
	router := gin.Default()
	{
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
		//swagger ui
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.LoadHTMLGlob("templates/*")
		router.Static("/assets", "./assets")
		{
			router.GET("/", IndexPageHandlerFunc(&listenPort, &env))
			router.GET("/file", GetExcelExportHandlerFunc)
			router.POST("/file", PostJSONDataHandlerFunc)
		}
	}
	router.Run(fmt.Sprintf(":%d", listenPort))
}
