package go_struct

import "time"

type CAD_check_errors struct {
	Check_status_id int
	Error_code int
}

type LUT_cad_errors struct {
	Id int
	Func_name string
}
type Timestamp time.Time

type Cad_check_status struct {
	ID int
	Status_code *int
	Last_update Timestamp
	Path string
	Ref_num int
	System_code int
}


func (Cad_check_status) TableName() string {
	return "CAD_check_status"
}

func (CAD_check_errors) TableName() string {
	return "CAD_check_errors"
}

func (LUT_cad_errors) TableName() string {
	return "LUT_cad_errors"
}