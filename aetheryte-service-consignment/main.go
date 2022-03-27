package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/lukeajtodd/eorzea-global/aetheryte-service-consignment/proto/consignment"
	vessel "github.com/lukeajtodd/eorzea-global/aetheryte-service-vessel/proto/vessel"
	"go-micro.dev/v4"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Create the service
	service := micro.NewService(
		micro.Name("aetheryte.consignment.service"),
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

	consignmentCollection := client.Database("aetheryte").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}

	vesselClient := vessel.NewVesselService("aetheryte.vessel.service", service.Client())
	h := &handler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(service.Server(), h)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
