package main

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"time"
)

const  (
	fraudReportSheetNameSummary = "Summary"

)
var 	LocationShanghai, _ = time.LoadLocation("Asia/Shanghai")

func main() {
	year, month, day := time.Now().Date()
	startDate := time.Date(year, month-1, day, 0, 0, 0, 0, LocationShanghai)
	endDate := time.Date(year, month, day, 0, 0, 0, 0, LocationShanghai)
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", fraudReportSheetNameSummary)
	f.MergeCell(fraudReportSheetNameSummary, "A2", "H2")
	f.SetSheetRow(fraudReportSheetNameSummary, "A2", &[]interface{}{"This file is to be password protected"})
	f.SetSheetRow(fraudReportSheetNameSummary, "A4",&[]interface{}{"Sheet","Frequency","Content","Who is to share","Report to be marked to"})
	f.SetSheetRow(fraudReportSheetNameSummary, "A5",&[]interface{}{"Users","Monthly","All users having more than 01 claim within 03  months","Igloo", ""})
	f.SetSheetRow(fraudReportSheetNameSummary, "A6",&[]interface{}{"Locations","Monthly","Top 20 area locations in term of claim frequency","Igloo", ""})
	f.SetSheetRow(fraudReportSheetNameSummary, "A7",&[]interface{}{"Repair Shops","Monthly","Top 10 repair shops in term of claim frequency","Igloo", ""})
	f.SetSheetRow(fraudReportSheetNameSummary, "A9", &[]interface{}{"Date From", startDate.Format("2006-01-02")})
	f.SetSheetRow(fraudReportSheetNameSummary, "A10", &[]interface{}{"Date To", endDate.Format("2006-01-02")})
	f.SetSheetRow(fraudReportSheetNameSummary, "A12", &[]interface{}{"Summary"})
	f.SetSheetRow(fraudReportSheetNameSummary, "A13", &[]interface{}{"1. All users having more than 01 claim within 03  months"})
	f.Path = "test.xlsx"
	f.Save()

}
