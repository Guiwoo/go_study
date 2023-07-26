package main

import (
	"db_design/table"
	"fmt"
)

func main() {
	db := GetDB()
	var survey table.Survey
	err := db.Model(&survey).Preload("Creator").Preload("Contents").Find(&survey, 1).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("survey is %+v", survey)
}
