package modals

import (
	"fmt"
	. "github.com/ahmetb/go-linq"
	//xlst "github.com/ivahaev/go-xlsx-templater"
	xlst "gitee.com/lucky51/go-xlsx-templater"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

// BasicInfo 住户基本信息
type BasicInfo struct {
	CustomerID   int    `json:"customerId"`
	CustomerName string `json:"customerName"`
	Gender       string `json:"gender"`
	Sex          string `json:"sex"`
	RoomCode     string `json:"roomCode"`
	BedCode      string `json:"bedCode"`
	CareLevel    string `json:"careLevel"`
	State        string `json:"state"`
	BirthOfDate  string `json:"birthOfDate"`
	Age          int    `json:"age"`
	InStartDate  string `json:"inStartDate"`
	InCode       string `json:"inCode"`
	StartDate    string `json:"startDate"`
	InDate       string `json:"inDate"`
}

// RecordItem 服务项
type RecordItem struct {
	GoodsInfoId   int    `json:"goodsInfoId"`
	Year          int    `json:"year"`
	Week          int    `json:"week"`
	StrWeek       string `json:"strweek"`
	ExecDate      string `json:"execDate"`
	FormatDate    string `json:"formatDate"`
	ServicePeriod string `json:"servicePeriod"`
	ServiceName   string `json:"serviceName"`
	Signature     string `json:"signature"`
	Reason        string `json:"reason"`
	CycleType     int8   `json:"cycleType"`
	ServiceState  int8   `json:"serviceState"`
	RecordType    int8   `json:"recordType"`
}

// OuterItem CMP接口返回数据项
type OuterItem struct {
	Year       int          `json:"year"`
	Week       int          `json:"week"`
	Date1      string       `json:"date1"`
	Date2      string       `json:"date2"`
	Date3      string       `json:"date3"`
	Date4      string       `json:"date4"`
	Date5      string       `json:"date5"`
	Date6      string       `json:"date6"`
	Date7      string       `json:"date7"`
	DayItems   []RecordItem `json:"dayItems"`
	WeekItems  []RecordItem `json:"weekItems"`
	MonthItems []RecordItem `json:"monthItems"`
	TempItems  []RecordItem `json:"tempItems"`
	UnItems    []RecordItem `json:"unItems"`
}

// OuterInput CMP接口返回数据
type OuterInput struct {
	Items     []OuterItem `json:"items"`
	BasicInfo BasicInfo   `json:"basicInfo"`
	SheetName string      `json:"sheetName"`
}

var firstCarePlanStartTime string = ""

// MapTo 接口数据转换
func MapTo(outers *[]OuterInput) []xlst.InParam {
	i:=0
	//查询有效的sheet集合，会出现items长度为0的数据
	arrLen :=From(*outers).CountWithT(func(o OuterInput) bool { return len(o.Items)>0})
	result := make([]xlst.InParam, arrLen)
	for key, item := range *outers {
		if key == 0 {
			firstCarePlanStartTime = item.BasicInfo.StartDate
		}
		if len(item.Items) ==0{
			continue
		}
		tempCtx, _ := item.GetRenderContext()
		result[i] = xlst.InParam{
			In: map[string]interface{}{
				"outer": tempCtx,
			},
			SheetName: item.SheetName,
		}
		i =i+1
	}
	return result
}

// isLessInDate 用于判断是否小于开始日期
func isLessInDate(year int, dStr string, inDateStr string) bool {
	t, err := time.ParseInLocation("2006-01-02", fmt.Sprintf("%d-%s", year, dStr), time.Local)
	if err != nil {
		zap.L().Error("isLessInDate parse error", zap.Error(err))
		return false
	}
	inDate, err := time.ParseInLocation("2006-01-02T15:04:05", inDateStr, time.Local)
	if err != nil {
		zap.L().Error("isLessInDate parse error", zap.Error(err))
		return false
	}
	return t.Before(inDate)
}

//GetRenderContext 返回用于模板渲染的字典
func (input *OuterInput) GetRenderContext() ([]map[string]interface{}, error) {
	var count int = 0
	if count = len(input.Items); count == 0 {
		count = 1
	}
	result := make([]map[string]interface{}, count)
	if len(input.Items) == 0 {
		result[0] = map[string]interface{}{
			"name":      input.BasicInfo.CustomerName,
			"bedCode":   input.BasicInfo.CustomerName,
			"careLevel": input.BasicInfo.CareLevel,
			"sex":       input.BasicInfo.Sex,
			"age":       strconv.Itoa(input.BasicInfo.Age), //fmt.Sprintf("%d",input.BasicInfo.Age)
			"inDate":    input.BasicInfo.InDate,
			"inCode":    input.BasicInfo.InCode,
			"dayServices": []map[string]interface{}{
				{
					"items": []map[string]interface{}{},
				},
			},
			"weekServices": []map[string]interface{}{
				{
					"items": []map[string]interface{}{},
				},
			},
			"monthServices": []map[string]interface{}{
				{
					"items": []map[string]interface{}{},
				},
			},
			"tempServices": []map[string]interface{}{
				{
					"items": []map[string]interface{}{},
				},
			},
			"unServices": []map[string]interface{}{
				{
					"items": []map[string]interface{}{},
				},
			},
		}

		defaultDict := result[0]
		dtS := strings.Split(input.BasicInfo.StartDate, "T")
		if len(dtS) > 0 {
			defaultDict["startDate"] = dtS[0]
		} else {
			defaultDict["startDate"] = ""
		}
	}

	for outerIdx, outer := range input.Items {
		//outer
		outerDict := make(map[string]interface{})
		outerDict["name"] = input.BasicInfo.CustomerName
		outerDict["bedCode"] = input.BasicInfo.BedCode
		outerDict["careLevel"] = input.BasicInfo.CareLevel
		outerDict["sex"] = input.BasicInfo.Sex
		outerDict["yearWeek"] = fmt.Sprintf("%d-%d", outer.Year, outer.Week)
		outerDict["age"] = fmt.Sprintf("%d", input.BasicInfo.Age)
		outerDict["inDate"] = input.BasicInfo.InDate
		dtS := strings.Split(input.BasicInfo.StartDate, "T")
		if len(dtS) > 0 {
			outerDict["startDate"] = dtS[0]
		} else {
			outerDict["startDate"] = ""
		}

		outerDict["inCode"] = input.BasicInfo.InCode
		//day
		dayService := []map[string]interface{}{
			{
				"items": make([]map[string]interface{}, len(outer.DayItems)),
			},
		}
		weekDay := [...]string{outer.Date1, outer.Date2, outer.Date3, outer.Date4, outer.Date5, outer.Date6, outer.Date7}
		for idx, day := range weekDay {
			dateK := fmt.Sprintf("date%d", idx+1)
			outerDict[dateK] = day
		}
		//dayService[0]["items"] = make([]map[string]interface{}, len(outer.DayItems))
		//dayDictItems := dayService[0]["items"].([]map[string]interface{})
		//fmt.Println("day item length :", len(outer.DayItems))
		//for wIdx, wVal := range weekDay {
		//	for sIdx, sItem := range outer.DayItems {
		//
		//		nameK := fmt.Sprintf("name%d", wIdx+1)
		//		if dayDictItems[sIdx] == nil {
		//			dayDictItems[sIdx] = make(map[string]interface{})
		//		}
		//		if sItem.FormatDate == wVal {
		//			dayDictItems[sIdx][nameK] = sItem.Signature
		//		} else {
		//			if isLessInDate(outer.Year, wVal, firstCarePlanStartTime) {
		//				dayDictItems[sIdx][nameK] = " "
		//			} else {
		//				dayDictItems[sIdx][nameK] = "--"
		//			}
		//
		//			dayDictItems[sIdx]["servicePeriod"] = sItem.ServicePeriod
		//			dayDictItems[sIdx]["serviceName"] = sItem.ServiceName
		//		}
		//	}
		//}
		// 1. 返回的数据按照 时间 和 服务名称 进行分组
		// 2. 遍历分组项， 循环每周的日期 比对，填充
		q := From(outer.DayItems).GroupBy(func(i interface{}) interface{} {
			return struct {
				Time        string
				ServiceName string
				GoodsInfoId int
			}{
				Time:        i.(RecordItem).ServicePeriod,
				ServiceName: i.(RecordItem).ServiceName,
				GoodsInfoId : i.(RecordItem).GoodsInfoId,
			}
		}, func(r interface{}) interface{} {
			return r
		}).OrderBy(func(g interface{}) interface{} {
			return g.(Group).Key.(struct {
				Time        string
				ServiceName string
				GoodsInfoId int
			}).Time
		})
		qResults := q.Results()
		dayService[0]["items"] = make([]map[string]interface{}, len(qResults))
		dayDictItems := dayService[0]["items"].([]map[string]interface{})
		for gpIdx, gpItem := range qResults {
			if dayDictItems[gpIdx] == nil {
				dayDictItems[gpIdx] = make(map[string]interface{})
			}
			currentGp := gpItem.(Group)
			for _, currentItem := range currentGp.Group {
				for idx, val := range weekDay {
					nameK := fmt.Sprintf("name%d", idx+1)
					temp := currentItem.(RecordItem)
					if val == temp.FormatDate {
						if temp.ServiceState ==2{
							dayDictItems[gpIdx][nameK] =fmt.Sprintf("__%s__", temp.Signature)
						}else{
							dayDictItems[gpIdx][nameK] =temp.Signature  //fmt.Sprintf("__%s__", temp.Signature)
						}
					} else {
						if _, ok := dayDictItems[gpIdx][nameK]; !ok {
							if isLessInDate(outer.Year,val,firstCarePlanStartTime){
								dayDictItems[gpIdx][nameK] = " "
							}else{
								dayDictItems[gpIdx][nameK] = "--"
							}
						}
					}
				}
			}

			dayDictItems[gpIdx]["servicePeriod"] = currentGp.Key.(struct {
				Time        string
				ServiceName string
				GoodsInfoId int
			}).Time
			dayDictItems[gpIdx]["serviceName"] = currentGp.Key.(struct {
				Time        string
				ServiceName string
				GoodsInfoId int
			}).ServiceName
		}

		//dayService[0]["items"] = make([]map[string]interface{}, len(outer.DayItems))
		//for dayIdx, day := range outer.DayItems {
		//	dayDictItems := dayService[0]["items"].([]map[string]interface{})
		//	if dayDictItems[dayIdx] == nil {
		//		dayDictItems[dayIdx] = make(map[string]interface{})
		//	}
		//	for idx, val := range weekDay {
		//		nameK := fmt.Sprintf("name%d", idx+1)
		//		if val == day.FormatDate {
		//			dayDictItems[dayIdx][nameK] = day.Signature //fmt.Sprintf("__%s__", day.Signature)
		//		} else {
		//			dayDictItems[dayIdx][nameK] = "--"
		//		}
		//	}
		//	dayDictItems[dayIdx]["servicePeriod"] = day.ServicePeriod
		//	dayDictItems[dayIdx]["serviceName"] = day.ServiceName
		//}

		outerDict["dayServices"] = dayService
		//week
		weekService := []map[string]interface{}{
			{},
		}
		weekServiceSlice := make([]map[string]interface{}, len(outer.WeekItems))
		for mIdx, m := range outer.WeekItems {
			weekDict := make(map[string]interface{})
			weekDict["formatDate"] = m.FormatDate
			weekDict["serviceName"] = m.ServiceName
			weekDict["signature"] = m.Signature
			if weekServiceSlice[mIdx] == nil {
				weekServiceSlice[mIdx] = make(map[string]interface{})
			}
			weekServiceSlice[mIdx] = weekDict
		}
		weekService[0]["items"] = weekServiceSlice
		outerDict["weekServices"] = weekService
		monthService := []map[string]interface{}{
			{},
		}
		monthServiceSlice := make([]map[string]interface{}, len(outer.MonthItems))
		for mIdx, m := range outer.MonthItems {
			monthDict := make(map[string]interface{})
			monthDict["formatDate"] = m.FormatDate
			monthDict["serviceName"] = m.ServiceName
			monthDict["signature"] = m.Signature
			if monthServiceSlice[mIdx] == nil {
				monthServiceSlice[mIdx] = make(map[string]interface{})
			}
			monthServiceSlice[mIdx] = monthDict
		}
		monthService[0]["items"] = monthServiceSlice
		outerDict["monthServices"] = monthService

		//临时的
		tempService := make([]map[string]interface{}, 1, 1)
		tempService[0] = make(map[string]interface{})
		tempServiceSlice := make([]map[string]interface{}, len(outer.TempItems))
		for tempIdx, tmp := range outer.TempItems {
			tempDict := make(map[string]interface{})
			tempDict["formatDate"] = tmp.FormatDate
			tempDict["serviceName"] = tmp.ServiceName
			tempDict["signature"] = tmp.Signature
			tempDict["servicePeriod"] = tmp.ServicePeriod
			if tempServiceSlice[tempIdx] == nil {
				tempServiceSlice[tempIdx] = make(map[string]interface{})
			}
			tempServiceSlice[tempIdx] = tempDict
		}
		tempService[0]["items"] = tempServiceSlice
		outerDict["tempServices"] = tempService
		//未执行的
		unExeService := []map[string]interface{}{
			{},
		}
		unExeServiceSlice := make([]map[string]interface{}, len(outer.UnItems))
		for unIdx, tmp := range outer.UnItems {
			unDict := make(map[string]interface{})
			unDict["formatDate"] = tmp.FormatDate
			unDict["serviceName"] = tmp.ServiceName
			unDict["reason"] = tmp.Reason
			unDict["servicePeriod"] = tmp.ServicePeriod
			if unExeServiceSlice[unIdx] == nil {
				unExeServiceSlice[unIdx] = make(map[string]interface{})
			}
			unExeServiceSlice[unIdx] = unDict
		}
		unExeService[0]["items"] = unExeServiceSlice
		outerDict["unServices"] = unExeService

		result[outerIdx] = outerDict
	}
	return result, nil
}
