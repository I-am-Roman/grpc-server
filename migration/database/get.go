package database

import (
	"fmt"
)

type User struct {
	id         string
	age        string
	first_name string
	last_name  string
	email      string
}

func (c *Connection) Get_all_users() []string {

	rows, err := c.connect.Query("select * from users")
	if err != nil {
		panic(err)
	}
	defer c.connect.Close()
	defer rows.Close()

	var one string
	var two string
	var three string
	var four string
	var five string

	products := []string{}
	for rows.Next() {
		fmt.Println("row - ", rows)

		err := rows.Scan(&one, &two, &three, &four, &five)
		if err != nil {
			fmt.Println(err)
			continue
		}
		res := one + " " + two + " " + three + " " + four + " " + five + " "
		products = append(products, res)

	}

	return products

	// p := User{}
	// users := []string{}
	// for rows.Next() {
	// 	fmt.Println("row - ", rows)

	// 	err := rows.Scan(&p.id, &p.age, &p.first_name, &p.last_name, &p.email)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	users = append(users)
	// }
	// return users

}
