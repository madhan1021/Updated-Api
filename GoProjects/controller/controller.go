package controller

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"math/rand"
	"net/http"

	"log"
	"os"

	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Madhan/GoProjects/models"
	"github.com/gorilla/mux"
)

var employees []models.Employees

//To get an employee
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func GetSummary(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	file, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Println(err)
	}
	var data models.Conf
	new_data, _ := ioutil.ReadAll(file)
	err2 := json.Unmarshal(new_data, &data)
	if err2 != nil {
		fmt.Println("error", err2)
	}
	json.NewEncoder(w).Encode(data.Summary)
}

// To delete employee
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id_param := mux.Vars(r)
	for index, item := range employees {
		if item.ID == id_param["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(item)
	}

}

//to get selected Employee
func Get1Employee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id_param := mux.Vars(r)
	for _, item := range employees {
		if item.ID == id_param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//To create new Employee with validation

func AddEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	new_emp := models.Employees{}
	json.NewDecoder(r.Body).Decode(&new_emp)
	valid := validator.New()
	err1 := valid.Struct(new_emp)
	if err1 != nil {
		validErrors := err1.(validator.ValidationErrors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error in creation of Employee")
		bodyResponse := map[string]string{"errors": validErrors.Error()}
		json.NewEncoder(w).Encode(bodyResponse)
		return
	}

	file, err := os.Open("./config/expenditure.json")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	newArray, _ := ioutil.ReadAll(file)
	var new_conf models.Expenditure
	err2 := json.Unmarshal(newArray, &new_conf)
	if err2 != nil {
		fmt.Println("error: ", err2.Error())
	}

	new_emp.ID = strconv.Itoa(rand.Intn(10000))
	new_emp.Savings = new_emp.CompanyInformation.Salary - (new_conf.Rent + new_conf.Bills + new_conf.Food)
	employees = append(employees, new_emp)
	fmt.Println()
	fmt.Println(new_emp.Email)
	fmt.Println(new_emp.Phone)
	// w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	// f := excelize.NewFile()
	// userFields := map[string]string{"A1": "Name", "B1": "Phone", "C1": "Email", "D1": "Salary"}
	// userValues := map[string]string{"A2": new_emp.Name, "B2": new_emp.Phone, "C2": new_emp.Email, "D2": string(rune(new_emp.Savings))}
	// for k, v := range userFields {
	// 	f.SetCellValue("Sheet1", k, v)
	// }
	// for k, v := range userValues {
	// 	f.SetCellValue("Sheet1", k, v)
	// }
	// err = f.SaveAs("newFile.xlsx")
	// if err != nil {
	// 	log.Fatalln("error: ", err.Error())
	// } else {
	// 	fmt.Println("Successfully created")
	// }

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(new_emp)
	// json.NewEncoder(w).Encode(map[string]string{"Name": new_emp.Name, "Phone": new_emp.Phone, "Email": new_emp.Email, "Salary": new_emp.ID})
	fmt.Println("User successfully added")

}

//To update the Employee

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id_param := mux.Vars(r)
	for index, item := range employees {
		if item.ID == id_param["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			var new_emp models.Employees
			_ = json.NewDecoder(r.Body).Decode(&new_emp)
			new_emp.ID = id_param["id"]
			employees = append(employees, new_emp)
			json.NewEncoder(w).Encode(new_emp)
			fmt.Printf("User----%v updated successfully\n", item.Name)
			return
		}

	}

}

// for excel part
func InsertExcel(w http.ResponseWriter, r *http.Request) {
	var new_emp models.Employees
	f := excelize.NewFile()
	userFields := map[string]string{"A1": "Name", "A2": "Phone", "A3": "Email", "A4": "Salary"}
	userValues := map[string]string{"B1": new_emp.Name, "B2": new_emp.Phone, "B3": new_emp.Email, "B4": string(rune(new_emp.Savings))}
	for k, v := range userFields {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range userValues {
		f.SetCellValue("Sheet1", k, v)
	}
	err := f.SaveAs("2newFile.xlsx")
	if err != nil {
		log.Fatalln("error: ", err.Error())
	} else {
		fmt.Println("Successfully created")
	}
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	json.NewEncoder(w).Encode(userValues)

}
