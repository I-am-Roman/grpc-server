package database

import "fmt"

func (c *Connection) Create_New_User(id string, age string, firts_name string, second_name string, email string) {

	defer c.connect.Close()

	sqlStatement := `
	  INSERT INTO users (id, age, first_name, last_name, email)
	  VALUES ($1, $2, $3, $4,$5)
	  RETURNING id`

	err := c.connect.QueryRow(sqlStatement, id, age, firts_name, second_name, email).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

}
