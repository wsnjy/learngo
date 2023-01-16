package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('budi', 'Budi')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Success  new customer data to table")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("id:", id)
		fmt.Println("name:", name)
	}

	defer rows.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)

		if err != nil {
			panic(err)
		}

		fmt.Println("====================")
		fmt.Println("id:", id, "name:", name, "email:", email, "balance:", balance, "rating:", rating, "birth_data:", birth_date, "married:", married, "created_at:", created_at)

		if email.Valid {
			fmt.Println(email.String)
		}

		if birth_date.Valid {
			fmt.Println(birth_date.Time)
		}
	}

	defer rows.Close()
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username from user where username = '" + username + "' and password = '" + password + "' limit 1"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		rows.Scan(&username)

		fmt.Println("Success login", username)
	} else {
		fmt.Printf("no authenticated")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username from user where username = ? and password = ? limit 1"
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		rows.Scan(&username)

		fmt.Println("Success login", username)
	} else {
		fmt.Printf("no authenticated")
	}
}

func TestExecSQLSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, "jering", "Jering")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Success  new customer data to table")
}

func TestExecSQLParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "bobi"
	password := "bikul"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Success  insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "test@test.com"
	comment := "lorem ipsum bla bla bla"

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("Last insert id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	stmt, err := db.PrepareContext(ctx, sqlQuery)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "bobi" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		lastInderId, _ := result.LastInsertId()

		fmt.Println("Comment id", lastInderId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	for i := 0; i < 10; i++ {
		email := "bobi" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, sqlQuery, email, comment)

		if err != nil {
			panic(err)
		}

		lastInderId, _ := result.LastInsertId()

		fmt.Println("Comment id", lastInderId)
	}

	err = tx.Commit()

	if err != nil {
		panic(err)
	}

}
