package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/lukeajtodd/eorzea-global/aetheryte-service-vessel/proto/vessel"
	"go-micro.dev/v4"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Create the service
	service := micro.NewService(
		micro.Name("aetheryte.vessel.service"),
	)

	// Initialise
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)

	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("aetheryte").Collection("vessels")

	repository := &MongoRepository{vesselCollection}

	h := &handler{repository}

	pb.RegisterVesselServiceHandler(service.Server(), h)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
