package helper

import (
	"context"
	"database/sql"
	"time"
)

func GetConnection2() *sql.DB {
	db, err := sql.Open("mysql", "...")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

type Presensi struct {
	JamMasuk  string
	JamPulang string
}

func GetOffice() Presensi {
	tx := GetConnection2()
	ctx := context.Background()
	script := "select jam_masuk, jam_pulang from office where id = ?"
	rows, err := tx.QueryContext(ctx, script, "1")
	PanicIfError(err)
	defer rows.Close()

	office := Presensi{}

	if rows.Next() {
		rows.Scan(&office.JamMasuk, &office.JamPulang)
		return office
	}
	return office
}
