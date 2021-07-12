package modals

import (
	"fmt"
	"testing"

	xlst "github.com/ivahaev/go-xlsx-templater"
)

func mapTo(outers *[]OuterInput) []xlst.InParam {
	result := make([]xlst.InParam, len(*outers))
	for key, item := range *outers {
		tempCtx, _ := item.GetRenderContext()
		result[key] = xlst.InParam{
			In: map[string]interface{}{
				"outer": tempCtx,
			},
			SheetName: item.SheetName,
		}
	}
	return result
}

// TestSheetsCreate 测试生成多sheets的表格
func TestSheetsCreate(t *testing.T) {
	items := []OuterInput{
		{
			Items: []OuterItem{
				{
					Date1: "05-04",
					Date2: "05-05",
					Date3: "05-06",
					Date4: "05-07",
					Date5: "05-08",
					Date6: "05-09",
					Date7: "05-10",
					DayItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					MonthItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					WeekItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					UnItems:   []RecordItem{},
					TempItems: []RecordItem{},
				},
				{
					Date1: "05-11",
					Date2: "05-12",
					Date3: "05-13",
					Date4: "05-14",
					Date5: "05-15",
					Date6: "05-16",
					Date7: "05-17",
					DayItems: []RecordItem{
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					MonthItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					WeekItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					UnItems: []RecordItem{
						{
							FormatDate:    "05-11",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "这个没有执行",
						},
					},
					TempItems: []RecordItem{
						{
							FormatDate:    "05-12",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
						},
					},
				},
				{
					Date1: "05-04",
					Date2: "05-05",
					Date3: "05-06",
					Date4: "05-07",
					Date5: "05-08",
					Date6: "05-09",
					Date7: "05-10",
					DayItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					MonthItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					WeekItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					UnItems:   []RecordItem{},
					TempItems: []RecordItem{},
				},
				{
					Date1: "05-11",
					Date2: "05-12",
					Date3: "05-13",
					Date4: "05-14",
					Date5: "05-15",
					Date6: "05-16",
					Date7: "05-17",
					DayItems: []RecordItem{
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					MonthItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					WeekItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					UnItems: []RecordItem{
						{
							FormatDate:    "05-11",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "这个没有执行",
						},
					},
					TempItems: []RecordItem{
						{
							FormatDate:    "05-12",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
						},
					},
				},
				{
					Date1: "05-11",
					Date2: "05-12",
					Date3: "05-13",
					Date4: "05-14",
					Date5: "05-15",
					Date6: "05-16",
					Date7: "05-17",
					DayItems: []RecordItem{
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					MonthItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					WeekItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					UnItems: []RecordItem{
						{
							FormatDate:    "05-11",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "这个没有执行",
						},
					},
					TempItems: []RecordItem{
						{
							FormatDate:    "05-12",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
						},
					},
				},
				{
					Date1: "05-11",
					Date2: "05-12",
					Date3: "05-13",
					Date4: "05-14",
					Date5: "05-15",
					Date6: "05-16",
					Date7: "05-17",
					DayItems: []RecordItem{
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					MonthItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					WeekItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					UnItems: []RecordItem{
						{
							FormatDate:    "05-11",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "这个没有执行",
						},
					},
					TempItems: []RecordItem{
						{
							FormatDate:    "05-12",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
						},
					},
				},
			},
			BasicInfo: BasicInfo{
				CustomerName: "张四",
				CustomerID:   123,
				BirthOfDate:  "06-06",
				InCode:       "123456",
				InDate:       "09-09",
				StartDate:    "1919-09-09",
				InStartDate:  "1919-09-09",
				RoomCode:     "1-201",
				BedCode:      "201",
				Age:          21,
				Sex:          "男",
				CareLevel:    "照护好几级",
			},
			SheetName: "20200504 1",
		},
		{
			Items: []OuterItem{
				{
					Date1: "09-07",
					Date2: "09-08",
					Date3: "09-09",
					Date4: "09-10",
					Date5: "09-11",
					Date6: "09-12",
					Date7: "09-13",
					DayItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					MonthItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					WeekItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					UnItems:   []RecordItem{},
					TempItems: []RecordItem{},
				},
				{
					Date1: "09-14",
					Date2: "05-15",
					Date3: "05-16",
					Date4: "05-17",
					Date5: "05-18",
					Date6: "05-19",
					Date7: "05-20",
					DayItems: []RecordItem{
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					MonthItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					WeekItems: []RecordItem{
						{
							FormatDate:    "05-05",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-09",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
						{
							FormatDate:    "05-10",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "",
						},
					},
					UnItems: []RecordItem{
						{
							FormatDate:    "05-11",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
							Reason:        "这个没有执行",
						},
					},
					TempItems: []RecordItem{
						{
							FormatDate:    "05-12",
							ServicePeriod: "2:00",
							ServiceName:   "服务名称1",
							Signature:     "张三",
						},
					},
				},
			},
			BasicInfo: BasicInfo{
				CustomerName: "张四",
				CustomerID:   123,
				BirthOfDate:  "06-06",
				InCode:       "123456",
				InDate:       "09-09",
				StartDate:    "1919-09-09",
				InStartDate:  "1919-09-09",
				RoomCode:     "1-201",
				BedCode:      "201",
				Age:          21,
				Sex:          "男",
				CareLevel:    "照护好几级",
			},
			SheetName: "20200907 2",
		},
	}
	ctxList := mapTo(&items)
	doc := xlst.New()
	err := doc.ReadTemplate("../excel/template-new.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	doc.RenderMultipleSheet(nil, ctxList...)
	doc.Save("./output.xlsx")
}
