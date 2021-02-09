package go_struct

import "time"

type CAD_check_errors struct {
	Check_Status_Id int
	Error_Code int
}

type LUT_cad_errors struct {
	CombinedKey int
	Field7 string
}
type Timestamp time.Time

type Cad_check_status struct {
	Id int
	Status_Code *int
	Last_Update_Date Timestamp
	Path string
	Ref_Num int
	System_Code int
}


func (Cad_check_status) TableName() string {
	return "CAD_Check_Status"
}

func (CAD_check_errors) TableName() string {
	return "CAD_Check_Errors"
}

func (LUT_cad_errors) TableName() string {
	return "TABLES_data"
}
