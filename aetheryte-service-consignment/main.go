package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/lukeajtodd/eorzea-global/aetheryte-service-consignment/proto/consignment"
	vessel "github.com/lukeajtodd/eorzea-global/aetheryte-service-vessel/proto/vessel"
	"go-micro.dev/v4"
)

// 1. Matches the surface of our service and the proto definition
type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// 2. Dummy repository that simulates a data store
type Repository struct {
	consignments []*pb.Consignment
}

// 3. Creates a new consignment
//
// DOCS(functions)~ Detail around function parameters, method implementations
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

/*
 * 4. The service must IMPLEMENT the methods available on our proto surface.
 * We currently only have a create method in the repo as can be seen on the interface above.
 */
type consignmentService struct {
	repo         repository
	vesselClient vessel.VesselService
}

// 4. The actual end point for creating a consignment from the outside
// This is the only method defined on the surface of our service.
// Creating a consginment
func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselRes, err := s.vesselClient.FindAvailable(ctx, &vessel.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	if vesselRes == nil {
		return errors.New("error fetching vessel")
	}

	if err != nil {
		return err
	}

	req.VesselId = vesselRes.Vessel.Id

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// 5. The return value here matches the response defined in the proto definition
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	// Create the service
	service := micro.NewService(
		micro.Name("aetheryte.consignment.service"),
	)

	// Initialise
	service.Init()

	vesselClient := vessel.NewVesselService("aetheryte.vessel.service", service.Client())

	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo, vesselClient}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
