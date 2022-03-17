package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/lukeajtodd/eorzea-global/aetheryte-service-vessel/proto/vessel"
	"go-micro.dev/v4"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

// Find an available vessel
//
// If a vessel's max weight and capacity are above spec then we can return it
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("no vessel found for that specification")
}

type vesselService struct {
	repo Repository
}

func (s *vesselService) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	// Create the datastore of possible vessels
	vessels := []*pb.Vessel{
		{
			Id:        "vessel0",
			Name:      "Blackjack",
			MaxWeight: 10000,
			Capacity:  200,
		},
	}

	// Add them to the repository
	repo := &VesselRepository{vessels}

	// Create the vessel service
	service := micro.NewService(
		micro.Name("aetheryte.vessel.service"),
	)

	service.Init()

	// Register the vessel service handler from proto
	if err := pb.RegisterVesselServiceHandler(service.Server(), &vesselService{repo}); err != nil {
		log.Panic(err)
	}

	// Run the stuff
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
