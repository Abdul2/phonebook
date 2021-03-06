package main

// the underscore means that the package is not in the scope of this project
// i had to import the 'pg' through a 'go get' call 

import (

	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

//
var selectStatement = `select id, name from contact limit 10`


func main() {
	
	
		fmt.Printf("testing..\n")


		// if ssl not enabled on the db the code will fail 
		// Open will validate the arguments but will not create a connection
		// to validate the source exist do a Ping 
		//replace user 
		db, err := sql.Open("postgres", "user=abdul-rashidabdi dbname=phonebook host=localhost sslmode=disable")

		
		//in case of error print the erro and exit 
		checkError(err)

		// try a Ping
		
		err = db.Ping()
		
		if err != nil {
		fmt.Printf(err.Error())
		return
		} else{
			
			fmt.Printf ("Ping worked: data source exist \n ")
			}

		//when are you done with db oprations exit
		defer db.Close()

		//always anticipate error
		err = doSelect(db)
	
		checkError(err)

}


//doSelect a poniter to db.Open() object that we have obtained in main()
//  if all goes well it will return nil otherwise it will return an error 
func doSelect(db *sql.DB) error {
	

	//use the db.Open() to return a prepared statement to be used at later stage 
	stmt, err := db.Prepare(selectStatement)
	
	checkError(err)

	//Query executes a prepared query statement with the given arguments
	// and returns the query results as a *Rows. [this is direct from Go docs]

	rows, err := stmt.Query()
	
	checkError(err)
 
 	//once done with statement close it.
	defer stmt.Close()

	for rows.Next() {
		
		var id string
		var name string
		
		// to avoid caller having copies of values, pointers are used 
		err = rows.Scan(&id, &name)

		// first %v is replaced by value of id and second %v by value of name 
		fmt.Printf("firstname: %v lastname: %v \n", id, name);
	}

	return nil
}

func checkError(err error){
	
		if err != nil {
			panic(err)
		}
	
}





