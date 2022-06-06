package main

import (
	"context"
	"github/ibadi-id/ecommerce/config"
	"github/ibadi-id/ecommerce/ent"
	"github/ibadi-id/ecommerce/ent/migrate"
	"log"

	"entgo.io/ent/dialect"
	"github.com/go-sql-driver/mysql"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client, err := newClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}
	defer client.Close()
	createDBSchema(client)
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func newDSN() string {
	mc := mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
			"charset":   config.C.Database.Params.Charset,
			"loc":       config.C.Database.Params.Loc,
		},
	}

	return mc.FormatDSN()
}

func newClient() (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	d := newDSN()

	return ent.Open(dialect.MySQL, d, entOptions...)
}
