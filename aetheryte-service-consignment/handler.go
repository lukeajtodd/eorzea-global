package main

import (
	"context"
	"errors"

	pb "github.com/lukeajtodd/eorzea-global/aetheryte-service-consignment/proto/consignment"
	vessel "github.com/lukeajtodd/eorzea-global/aetheryte-service-vessel/proto/vessel"
)

type handler struct {
	repository
	vesselClient vessel.VesselService
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
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

	if err = s.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
