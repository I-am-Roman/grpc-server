package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	name string
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return nil
}

func (s *Service) Get(index int) (*Product, error) {
	return nil, nil
}

const (
	host     = "localhost"
	port     = 5432
	user     = "rtrashchinskiy"
	password = "your-password"
	dbname   = "my_database"
)

type Connection struct {
	connect *sql.DB
}

func NewConnection() *Connection {
	return &Connection{
		connect: plug_in(),
	}
}

// func (c *Connection) Get(index int) (*Product, error) {
// 	return nil, nil
// }

// func (c *Connection) post(index int) (*Product, error) {
// 	return nil, nil
// }

// func (c *Connection) Get(index int) (*Product, error) {
// 	return nil, nil
// }

func plug_in() *sql.DB {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recover from panic")
		}
	}()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succefully!")

	// CREATE TABLE users (
	// 	id SERIAL,
	// 	age INT,
	// 	first_name VARCHAR(255),
	// 	last_name VARCHAR(255),
	// 	email TEXT
	//   );

	return db
	// sqlStatement := `
	//   INSERT INTO users (id, age, first_name, last_name, email)
	//   VALUES ($1, $2, $3, $4,$5)
	//   RETURNING id`
	// id := 0
	// err = db.QueryRow(sqlStatement, id, 20, "Erorn", "Bush", "Erorn@main.ru").Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("New record ID is:", id)

	// rows, err := db.Query("select * from users")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	fmt.Println("row - ", rows)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}

	// }

}

// fmt.Println("Successfully connected!")
// tx.Prepare(`do $$ begin execute format($f$create table %I()$f$,$1); end; $$;`)
// db.Exec(dbname)

// CREATE TABLE users (
// 	id SERIAL PRIMARY KEY,
// 	age INT,
// 	first_name TEXT,
// 	last_name TEXT,
// 	email TEXT UNIQUE NOT NULL
//   );

//   insert := sq.Insert(conversationBaseName).
//   Columns(
// 	  "post_id",
// 	  "root_id",
// 	  "user_id",
// 	  "created_at",
// 	  "updated_at",
// 	  "group_id",
// 	  "chanel_id",
// 	  "duty_id",
//   ).
//   Values(
// 	  post.PostID,
// 	  post.RootID,
// 	  post.UserID,
// 	  time.Now().UTC(),
// 	  time.Now().UTC(),
// 	  post.GroupID,
// 	  post.ChanelID,
// 	  dbArray,
//   ).
//   Suffix("ON CONFLICT DO NOTHING").
//   PlaceholderFormat(sq.Dollar)

// _, err := db.Execx(ctx, insert)
// if err != nil {
// 	return err
// }

// fmt.Println()
// db, err := sql.Open("pgx", "postgres://password@localhost:5432/db")
// if err != nil {
// 	log.Fatal("sql.Open() - ", err)
// }
// ctx := context.Background()
// err = db.PingContext(ctx)
// if err != nil {
// 	log.Fatal("db.PingContext(ctx)", err)
// }

// fmt.Println("Done")

// type User struct {
// 	id         string
// 	age        string
// 	first_name string
// 	last_name  string
// 	email      string
// }

// func (c *Connection) get_all_users() []User {

// 	db := c.plug_in()
// 	rows, err := db.Query("select * from users")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	p := User{}
// 	users := []User{}
// 	for rows.Next() {
// 		fmt.Println("row - ", rows)

// 		err := rows.Scan(&p.id, &p.age, &p.first_name, &p.last_name, &p.email)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		users = append(users, p)
// 	}
// 	return users

// }
