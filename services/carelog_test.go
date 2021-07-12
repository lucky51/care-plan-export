package services

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCareLogRequest(t *testing.T) {
	res, err := CareResidentRecordGroupByWeek(&CareResidentRequestInput{
		TenantID:   9,
		ResidentID: 1389,
		StartDate:  "2014-01-01", //time.Date(2014,1,1,0,0,0,0,time.Local),
		EndDate:    "2020-09-17", //time.Date(2020,9,17,0,0,0,0,time.Local),
	})
	if err != nil {
		t.Error(err)
	} else {
		jsonRes, _ := json.Marshal(res.Result.Data)
		t.Log(string(jsonRes))
	}
}

func TestFmtTime(t *testing.T) {
	s1:= struct { 
		T time.Time
	}{
		time.Now(),
	}
	s, _ := json.Marshal(s1)
	s2:= struct{
		T time.Time
	}{
		
	}
	e:=json.Unmarshal([]byte("{\"T\":\"2021-01-06T12:17:27.68\"}"),&s2 )
	if e!=nil{
		fmt.Println(e)
	}
	//
	fmt.Println(string(s))
	fmt.Printf("%v \n",&s2)

	 t1,err :=time.Parse("2006-01-02T15:04:05.00","2021-01-06T12:17:27.68")
	 if err!=nil{
	 	fmt.Println(err)
	 }
	fmt.Printf("%v",t1)

}
func TestCompareTime(t *testing.T) {
	t1,err := time.Parse("2006-01-02","2009-01-01")
	if err!=nil{
		t.Fatal(err)
	}
	t2,err:=time.Parse("2006-01-02","2020-01-01")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log( t1.Before(t2) )
}