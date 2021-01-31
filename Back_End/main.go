package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


// Customer Struct(Model)
type Customer struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
	Address string `json:"address"`
	Nic 	string `json:"nic"`
	Contact int64 `json:"contact"`
}

//Init customer var as a slice customer struct
var customers [] Customer

// enable cross Origin policy
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}


// getAllCustomers
func findAllCustomers(w http.ResponseWriter, r *http.Request){
	var customers [] Customer
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type","application/json")

	db, _ := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/golangSimpleApplication")
	row, err := db.Query("select * from customer")
	if err != nil {
		panic(err.Error())
	}else{
		//var allCustomers [] Customer
		for row.Next(){
			var id string
			var name string
			var address string
			var nic string
			var contact int64
			err2 := row.Scan(&id, &name, &address, &nic, &contact)
			row.Columns()
			if err2 != nil{
				panic(err2.Error())
			}else{
				customer := Customer{
					ID:      id,
					Name:    name,
					Address: address,
					Nic:     nic,
					Contact:contact,
				}
				customers = append(customers, customer)
			}
		}
	}
	defer row.Close()
	json.NewEncoder(w).Encode(customers)

}

// getOneCustomer
func getCustomer(w http.ResponseWriter, r *http.Request)  {
	var customer1 Customer

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	db, _ := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/golangSimpleApplication")
	row, err := db.Query("select * from customer where id=?", params["id"])
	if err != nil {
		panic(err.Error())
	}else{
		//var allCustomers [] Customer
		for row.Next(){
			var id string
			var name string
			var address string
			var nic string
			var contact int64
			err2 := row.Scan(&id, &name, &address, &nic, &contact)
			row.Columns()
			if err2 != nil{
				panic(err2.Error())
			}else{
				customer := Customer{
					ID:      id,
					Name:    name,
					Address: address,
					Nic:     nic,
					Contact:contact,
				}
				//customers = append(customers, customer)
				customer1 = customer
			}
		}
	}
	defer row.Close()
	json.NewEncoder(w).Encode(customer1)

}

// saveCustomer
func saveCustomer(w http.ResponseWriter, r *http.Request)  {

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type","application/json")
	var customer Customer
	_= json.NewDecoder(r.Body).Decode(&customer)
	/*customer.ID = strconv.Itoa(rand.Intn(1000))
	customers = append(customers,customer)*/
	//json.NewEncoder(w).Encode(customer)


	db, _ := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/golangSimpleApplication")
	insert, err := db.Query("INSERT INTO customer (id, name, address, nic, contact) VALUES (?, ?, ?, ?, ?)", customer.ID, customer.Name, customer.Address, customer.Nic, customer.Contact)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	json.NewEncoder(w).Encode(customer)
}

// updateCustomer
func updateCustomer(w http.ResponseWriter, r *http.Request)  {

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	var customer Customer
	_= json.NewDecoder(r.Body).Decode(&customer)

	db, _ := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/golangSimpleApplication")
	update, err := db.Query("update customer set name=? , address=? , nic=? , contact=?  where id= ?", customer.Name, customer.Address, customer.Nic, customer.Contact, params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()
	json.NewEncoder(w).Encode(customer)
}

// deleteCustomer
func deleteCustomers(w http.ResponseWriter, r *http.Request)  {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	db, _ := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/golangSimpleApplication")
	delete, err := db.Query("delete from customer where id=?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
	json.NewEncoder(w).Encode(customers)
}


func main()  {

	// Init Router
	r := mux.NewRouter()

	//Route Handlers / Endpoint
	r.HandleFunc("/api/customer",findAllCustomers).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/customer/{id}",getCustomer).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/customer",saveCustomer).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/customer/{id}",updateCustomer).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/customer/{id}",deleteCustomers).Methods("DELETE", "OPTIONS")

	fmt.Println("server stat")
	log.Fatal(http.ListenAndServe(":8000", r))

}
