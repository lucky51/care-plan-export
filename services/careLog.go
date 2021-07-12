package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goexcel/modals"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// CareResidentRequestInput 照护执行计划查询接口
type CareResidentRequestInput struct {
	ResidentID int    `json:"residentId" form:"residentId" binding:"required"`
	TenantID   int    `json:"tenantId" form:"tenantId" binding:"required"`
	StartDate  string `json:"startDate" form:"startDate" binding:"required"`
	EndDate    string `json:"endDate" form:"endDate" binding:"required"`
}

// AbpAjaxResponseError ABP中的错误输出
type AbpAjaxResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	//	Details string	`json:"details"`
}

// ExPlanWrapResult 获取记录的外套结构，为了包含用于拼接文件名称的数据
type ExPlanWrapResult struct {
	ResidentName string              `json:"residentName"`
	Data         []modals.OuterInput `json:"data"`
}

// AbpAjaxResponse ABP返回的公共结构
type AbpAjaxResponse struct {
	Result              ExPlanWrapResult     `json:"result"`
	TargetURL           string               `json:"targetUrl"`
	Success             bool                 `json:"success"`
	Error               AbpAjaxResponseError `json:"error"`
	UnAuthorizedRequest bool                 `json:"unAuthorizedRequest"`
	Abp                 bool                 `json:"__abp"`
}

const careResidentRecordRequestSegmentURL string = "/api/services/app/CareLog/CareResidentRecordGroupByWeek"

// PingCmp 尝试请求CMP接口地址
func PingCmp() error {
	_, err := http.Head(viper.GetString("cmpUrl"))
	return err
}

//CareResidentRecordGroupByWeek 微服务调用，获取照护计划列表
func CareResidentRecordGroupByWeek(input *CareResidentRequestInput) (AbpAjaxResponse, error) {
	output := AbpAjaxResponse{}
	cmpBaseURL := viper.GetString("cmpUrl")

	if strings.HasSuffix(cmpBaseURL, "/") {
		cmpBaseURL = strings.TrimSuffix(cmpBaseURL, "/")
	}
	if strings.HasSuffix(cmpBaseURL, "\\") {
		cmpBaseURL = strings.TrimSuffix(cmpBaseURL, "\\")
	}
	requestURL := fmt.Sprintf("%s%s", cmpBaseURL, careResidentRecordRequestSegmentURL)
	fmt.Println(requestURL)
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(input)
	req, reqError := http.NewRequest("POST", requestURL, buf)
	if reqError != nil {
		return output, reqError
	}
	req.Header.Add("Oms-Tenant-Id", strconv.Itoa(input.TenantID))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		return output, e
	}
	defer res.Body.Close()
	fmt.Println(fmt.Sprintf("statu code : %d\n", res.StatusCode))
	decodeError := json.NewDecoder(res.Body).Decode(&output)
	if decodeError != nil {
		return output, decodeError
	}
	return output, nil
}
