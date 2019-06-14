package main

import (
	_"errors"
	"time"
)

// define structure similar to the query
type Student struct {
	ID        int
	name      string
	Age       int16
	Active    bool
	CreatedAT time.Time
	UpdatedAT time.Time
}

// select to the table students
func allstudents() (estudiantes []Student, err error){
	q := `select id, name, age, active, created_at from students`

	db := getConnection()
	defer db.Close()

	rows, err := db.Query(q)

	if err != nil{
		return
	}
	defer rows.Close()

	for rows.Next() {
		e := Student{}
		err = rows.Scan(
			&e.ID,
			&e.name,
			&e.Age,
			&e.Active,
			&e.CreatedAT,
		)

		if err != nil {
			return
		}
		estudiantes = append(estudiantes, e)
	}

	return estudiantes, nil
}



/*

func creation(e student) error{
	q := `INSERT INTO
 		  students (name, age, active)
          VALUES ($1, $2, $3)`
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)

	if err != nil{
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.name, e.Age, e.Active)
	if err != nil{
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1{
		return errors.New("idiota se espera una linea afectada")
	}
	return nil
}
*/
