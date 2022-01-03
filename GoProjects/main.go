package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Madhan/GoProjects/controller"
	"github.com/Madhan/GoProjects/models"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	var employees []models.Employees
	var new_emp models.Employees
	employees = append(employees, models.Employees{Name: "John Doe", Phone: "1234567890", Email: "abcdef@gmail.com", Technologies: []string{"golang", "java", "mysql"}, Savings: new_emp.Savings, CompanyInformation: &models.CompanyInformation{Company_Name: "NTT", Company_Location: "Bangalore", Position: "Software developer", Salary: 50000}, Address: &models.Address{City: "delhi", State: "UP", Country: "India"}})

	r.HandleFunc("/v1/user", controller.GetEmployee).Methods("GET")
	r.HandleFunc("/v1/user/{id}", controller.Get1Employee).Methods("GET")
	r.HandleFunc("/v1/user", controller.AddEmployee).Methods("POST")
	r.HandleFunc("/v1/user/{id}", controller.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/v1/user/{id}", controller.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/v1/summary", controller.GetSummary).Methods("GET")
	r.HandleFunc("/v1/user/excel", controller.InsertExcel).Methods("GET")
	//----------------------------------------------------------------------------------
	//to get the port value from config.json file

	file, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var data models.Conf
	new_data, _ := ioutil.ReadAll(file)
	err2 := json.Unmarshal(new_data, &data)
	if err2 != nil {
		fmt.Println("error", err2)
	}

	// //starting server
	fmt.Println("Starting Server....")
	log.Fatal(http.ListenAndServe(data.Port, r))

	// employees,err = http.Get("http://localhost:8080/v1/user")
	// if err != nil {
	// 	log.Fatal("error: ", err.Error())
	// } else {
	// 	data,_ := ioutil.ReadAll()
	// 	fmt.Println(string(employees))
	// }

}
