package CustomerRout

import (
	"BackEnd/DB"
	//Array and slice values encode as JSON arrays
	"encoding/json"
	//This is import to write values to the console

	// This implements a request router and dispatcher for matching  incoming requests to their
	"github.com/gorilla/mux"
	//This gives us HTTP client and server implementation for building the actual API
	"net/http"

	"encoding/base64"
)

var customers []customer //Create a Customer Array
// Customer struct
type customer struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Salary  float64 `json:"salary"`
}

/*------------------------------Save Customer----------------------------------------------------*/
func SaveCustomer(w http.ResponseWriter, r *http.Request) {
	//allow cross origin
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	id := base64.StdEncoding.EncodeToString([]byte(r.FormValue("id")))
	name := base64.StdEncoding.EncodeToString([]byte(r.FormValue("name")))
	address := base64.StdEncoding.EncodeToString([]byte(r.FormValue("address")))
	salary := r.FormValue("salary")

	//get connection
	db := DB.GetConnection()
	insert, err := db.Query("INSERT INTO customer VALUES (?, ?, ?, ?)", id, name, address, salary)
	if err != nil {
		panic(err.Error())
	}
	//close the executing
	defer insert.Close()
}

/*------------------------------Delete  Customer----------------------------------------------------*/
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	//allow cross origin
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	id := base64.StdEncoding.EncodeToString([]byte(r.FormValue("id")))
	//get connection
	db := DB.GetConnection()
	delete, err := db.Query("delete from customer where id=?", id)
	if err != nil {
		panic(err.Error())
	}
	//close the executing
	defer delete.Close()
	//return response
	json.NewEncoder(w).Encode(customers)

}

/*------------------------------Update  Customer----------------------------------------------------*/
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	//allow cross origin
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	id := base64.StdEncoding.EncodeToString([]byte(r.FormValue("id")))
	name := base64.StdEncoding.EncodeToString([]byte(r.FormValue("name")))
	address := base64.StdEncoding.EncodeToString([]byte(r.FormValue("address")))
	salary := r.FormValue("salary")

	//get connection
	db := DB.GetConnection()
	insert, err := db.Query("UPDATE customer set name=?,address=?,salary=? where id=?", name, address, salary, id)
	if err != nil {
		panic(err.Error())
	}
	//close the executing
	defer insert.Close()
}

/*------------------------------Get All  Customer----------------------------------------------------*/
func AllCustomer(w http.ResponseWriter, r *http.Request) {
	//create customer array
	var allCustomer []customer
	//accept cross origin
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	//set header content type application json
	w.Header().Set("Content-Type", "application/json")
	//get connection
	db := DB.GetConnection()
	//sql get all customer
	row, err := db.Query("Select * from customer")
	if err != nil {
		panic(err.Error())
	} else {
		//get all Customers in data base
		for row.Next() {
			var id string
			var name string
			var address string
			var salary float64

			err2 := row.Scan(&id, &name, &address, &salary)

			row.Columns()
			if err2 != nil {
				panic(err2.Error())
			} else {
				id1, _ := base64.StdEncoding.DecodeString(id)
				name1, _ := base64.StdEncoding.DecodeString(name)
				address1, _ := base64.StdEncoding.DecodeString(address)
				salary1 := salary

				// Initializing Customer struct
				custList := customer{
					ID:      string(id1),
					Name:    string(name1),
					Address: string(address1),
					Salary:  salary1,
				}
				//assign cuttList to allCustomer array
				allCustomer = append(allCustomer, custList)
			}
		}
	}
	//close the execution
	defer row.Close()
	//return response

	json.NewEncoder(w).Encode(allCustomer)

}

/*------------------------------Search  Customer----------------------------------*/
func SearchCustomer(w http.ResponseWriter, r *http.Request) {
	var cust customer

	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//Query strings in Go are accessed
	params := mux.Vars(r)
	//get connection
	db := DB.GetConnection()
	id := params["id"]
	id1 := base64.StdEncoding.EncodeToString([]byte(id))
	row, err := db.Query("select * from customer where id=?", id1)
	if err != nil {
		panic(err.Error())
	} else {
		for row.Next() {
			var id string
			var name string
			var address string
			var salary float64

			err2 := row.Scan(&id, &name, &address, &salary)
			row.Columns()
			if err2 != nil {
				panic(err2.Error())
			} else {
				// Initializing Customer struct
				id1, _ := base64.StdEncoding.DecodeString(id)
				name1, _ := base64.StdEncoding.DecodeString(name)
				address1, _ := base64.StdEncoding.DecodeString(address)
				salary1 := salary
				customer := customer{
					ID:      string(id1),
					Name:    string(name1),
					Address: string(address1),
					Salary:  salary1,
				}
				cust = customer
			}
		}
	}
	//close the executing
	defer row.Close()
	json.NewEncoder(w).Encode(cust)
}

/*--------------------------------setupCorsResponse------------------------------------------------*/
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) { //allow cross origin

	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	(*w).Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")

}
