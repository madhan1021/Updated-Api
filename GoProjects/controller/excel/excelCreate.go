package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f := excelize.NewFile()
	userFields := map[string]string{"A1": "Name", "A2": "Phone", "A3": "Email", "A4": "Salary"}
	userValues := map[string]string{"B1": "John", "B2": "8877888777", "B3": "abcdef@gmail.com", "B4": "25000"}
	userValues2 := map[string]string{"C1": "Dante", "C2": "8877888777", "C3": "abcdef@gmail.com", "C4": "25000"}
	for k, v := range userFields {
		f.SetCellValue("Sheet1", k, v)
	}

	for k, v := range userValues {
		f.SetCellValue("Sheet1", k, v)
	}

	for k, v := range userValues2 {
		f.SetCellValue("Sheet1", k, v)
	}

	
	err := f.SaveAs("4ExcelFile2.xlsx")
	if err != nil {
		log.Fatalln("error: ", err.Error())
	} else {
		fmt.Println("Successfully created")
	}

	readFile()
}

func readFile() {
	f, err := excelize.OpenFile("4ExcelFile2.xlsx")
	if err != nil {
		log.Fatalln("error: ", err.Error())
	}
	g := f.GetCellValue("Sheet1", "A2")
	if err != nil {
		log.Fatalln("error :", err.Error())
	}
	fmt.Println(g)
	rows := f.GetRows("Sheet1")
	for idx, row := range rows {
		if idx < 14 {
			continue
		}
		for _, colCell := range row {
			fmt.Println(colCell, "/t")
		}
		fmt.Println()
	}
	// g1 := f.GetCellValue("Sheet1", "A1")
	// if err != nil {
	// 	log.Fatalln("error :", err.Error())
	// }
	// fmt.Println(g1)
	// g2 := f.GetCellValue("Sheet1", "A3")
	// if err != nil {
	// 	log.Fatalln("error :", err.Error())
	// }
	// fmt.Println(g2)

	// g3 := f.GetCellValue("Sheet1", "B3")
	// if err != nil {
	// 	log.Fatalln("error :", err.Error())
	// }
	// fmt.Println(g3)

}
