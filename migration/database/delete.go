package database

func (c *Connection) Delete_User(email string) {

	defer c.connect.Close()

	_, err := c.connect.Exec("delete from users where email = $1", email)
	if err != nil {
		panic(err)
	}

}
