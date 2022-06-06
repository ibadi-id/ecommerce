package main

import (
	"github/ibadi-id/ecommerce/config"
	"github/ibadi-id/ecommerce/ent"
	"github/ibadi-id/ecommerce/pkg/adapter/controller"
	"github/ibadi-id/ecommerce/pkg/infrastructure/datastore"
	"github/ibadi-id/ecommerce/pkg/infrastructure/graphql"
	"github/ibadi-id/ecommerce/pkg/infrastructure/router"
	"github/ibadi-id/ecommerce/pkg/registry"
	"log"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	ctrl := newController(client)

	srv := graphql.NewServer(client, ctrl)
	e := router.New(srv)

	log.Println(config.C.Server.Address)
	e.Logger.Fatal(e.Start(":" + config.C.Server.Address))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}

	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}
