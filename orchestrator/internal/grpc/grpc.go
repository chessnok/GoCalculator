package grpc

import (
	"context"
	"fmt"
	"github.com/chessnok/GoCalculator/orchestrator/internal/db"
	pb "github.com/chessnok/GoCalculator/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	pb.OrchestratorServiceServer
	calculatorConfig *pb.Config
	db               *db.Postgres
}

func NewServer(calculatorConfig *pb.Config, agents *db.Postgres) *Server {
	return &Server{
		calculatorConfig: calculatorConfig,
		db:               agents,
	}
}

func (s *Server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	s.db.Agents.SetAgentLastPing(req.AgentId)
	// TODO: add saving last expression
	return &pb.PingResponse{
		NewConfig: s.calculatorConfig,
	}, nil
}
func newAgentId() string {
	uid := uuid.New().String()
	return uid
}
func (s *Server) RegisterAgent(ctx context.Context, request *pb.RegisterAgentRequest) (*pb.RegisterAgentResponse, error) {
	uid := newAgentId()
	_ = s.db.Agents.NewAgent(uid, request.Pid)
	return &pb.RegisterAgentResponse{
		AgentId: uid,
		Config:  s.calculatorConfig,
	}, nil
}

func Start(calculatorConfig *pb.Config, db *db.Postgres) error {
	host := "localhost"
	port := "5000"

	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	orchestratorServer := NewServer(calculatorConfig, db)
	pb.RegisterOrchestratorServiceServer(grpcServer, orchestratorServer)
	return grpcServer.Serve(lis)
}
