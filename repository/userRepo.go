package repository

import (
	"context"
	"go-pgx-loco/db"
	"go-pgx-loco/model"
	"go-pgx-loco/util"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func AddUser(u *model.User) error {
	defer util.Timer(time.Now(), "register")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	dob, err := time.Parse("2006-01-02", u.DOB)
	if err != nil {
		return err
	}

	conn, err := db.InitDB()
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO USERS (username,fullname,password,email,created_at,dob) values($1,$2,$3,$4,$5,$6)", u.Username, u.Fullname, hashedPassword, u.Email, time.Now(), dob)
	if err != nil {
		go db.CloseDB(conn)
		return err
	}
	go db.CloseDB(conn)
	return nil
}
