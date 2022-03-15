package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/lukeajtodd/aetheryte-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// 1. Matches the surface of our service and the proto definition
type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// 2. Dummy repository that simulates a data store
type Repository struct {
	mu           sync.RWMutex // DOCS(locks)~
	consignments []*pb.Consignment
}

// 3. Creates a new consignment
//
// DOCS(functions)~ Detail around function parameters, method implementations
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
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
type service struct {
	repo repository
}

// 4. The actual end point for creating a consignment from the outside
// This is the only method defined on the surface of our service.
// Creating a consginment
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// 5. The return value here matches the response defined in the proto definition
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &Repository{}

	// Setup the gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	s := grpc.NewServer()

	// Register our service with the server
	// Our implementation gets tied into the auto-generated interface from the proto definition
	//
	// DOCS(gRPC)~
	pb.RegisterShippingServiceServer(s, &service{repo})

	// Register reflection on server
	//
	// DOCS(gRPC)~
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
