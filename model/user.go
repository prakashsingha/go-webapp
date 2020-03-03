package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/prakashsingha/go-webapp/helper"
)

const passwordSalt = "mPRNYu0y6fuaImTNYyv0ZkwKtfwpM-oTVIKfGKMAJbNJ4Wu3mpmTHuyUC95jH-vm2x8Gqy-yoRbY3GgqTczGzw"

type User struct {
	ID        string
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

func Login(email, password string) (*User, error) {
	result := &User{}
	pwd := helper.Hasher(passwordSalt, email, password)

	row := db.QueryRow(`
					SELECT 
						id, 
						email, 
						firstname, 
						lastname,
						lastlogin
					FROM 
						public.user
					WHERE 
						email = $1
					AND
						password = $2`, email, pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName, &result.LastLogin)

	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}

	//update LastLogin field
	t := time.Now()
	fmt.Println("Time: ", t)
	_, err = db.Exec(`
					UPDATE 
						public.user
					SET
						lastlogin = $1
					WHERE
						id = $2`, t.Format("2006-01-02 15:04:05"), result.ID)
	if err != nil {
		log.Printf("Failed to update login time for user %v to %v: %v", result.Email, t, err)
	}

	return result, nil

}
