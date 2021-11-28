package main

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/x64fun/terra/example/pb/terra/wxwork"

	// postgreSQL 驱动
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	dsn := "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=postgres sslmode=disable"
	db, _ := sqlx.Connect("postgres", dsn)
	tx := db.MustBegin()
	wxwork.DEFAULT_TABLE_PREFIX = `"scrm_wxwork"`
	list, err := wxwork.GetDepartmentAllList(ctx, tx, nil, wxwork.DepartmentTableColumnID[tx.DriverName()], wxwork.DepartmentTableColumnName[tx.DriverName()])
	if err != nil {
		log.Println(err)
		return
	}
	for _, item := range list {
		log.Println(item.ID, item.Name)
	}
}
