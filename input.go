package input

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func input() {
	var number, apartment int
	var name, phone_number, email, city, district, street string
	connStr := "user=xyz password=1111 dbname=apelsin sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Print("Input name: ")
	fmt.Scan(&name)
	fmt.Print("Input phone number: ")
	fmt.Scan(&phone_number)
	fmt.Print("Input email: ")
	fmt.Scan(&email)
	fmt.Print("Input city: ")
	fmt.Scan(&city)
	fmt.Print("Input district: ")
	fmt.Scan(&district)
	fmt.Print("Input street: ")
	fmt.Scan(&street)
	fmt.Print("Input number: ")
	fmt.Scan(&number)
	fmt.Print("Input apartment: ")
	fmt.Scan(&apartment)

	result, err := db.Exec("insert into customers(name, phone_number, email, city, district, street, number, apartment) values ( $1, $2, $3, $4, $5 ,$6, $7, $8)",
		&name, &phone_number, &email, &city, &district, &street, &number, &apartment)
	fmt.Println(result.RowsAffected())
}
