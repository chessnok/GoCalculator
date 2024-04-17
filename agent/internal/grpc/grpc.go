package grpc

import (
	"context"
	pb "github.com/chessnok/GoCalculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strconv"
	"time"
)

type GRPC struct {
	cfg              *Config
	conn             *grpc.ClientConn
	calculatorConfig *pb.Config
	agentId          string
}

func NewGRPC(cfg *Config) *GRPC {
	return &GRPC{
		cfg: cfg,
	}
}
func (g *GRPC) startPing() {
	for {
		err := g.ping()
		if err != nil {
			log.Println("ping failed: ", err)
		}
		time.Sleep(1 * time.Second)
	}

}
func (g *GRPC) Connect() error {
	addr := g.cfg.Address()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	g.conn = conn
	err = g.RegisterAgent()
	if err != nil {
		g.conn.Close()
		return err
	}
	go g.startPing()
	return nil
}

func (g *GRPC) Close() {
	g.conn.Close()
}

func (g *GRPC) RegisterAgent() error {
	client := pb.NewOrchestratorServiceClient(g.conn)
	ctx := context.Background()
	resp, err := client.RegisterAgent(ctx, &pb.RegisterAgentRequest{
		Pid: strconv.Itoa(os.Getpid()),
	})
	if err != nil {
		return err
	}
	g.agentId = resp.AgentId
	return nil
}

func (g *GRPC) ping() error {
	client := pb.NewOrchestratorServiceClient(g.conn)
	ctx := context.Background()
	resp, err := client.Ping(ctx, &pb.PingRequest{
		LastTask: "implement me",
		AgentId:  g.agentId,
	})
	if err != nil {
		return err
	}
	g.calculatorConfig = resp.NewConfig
	return nil
}

//// закроем соединение, когда выйдем из функции
//defer conn.Close()
//grpcClient := pb.NewOrchestratorServiceClient(conn)
//ctx := context.Background()
//resp, err := grpcClient.RegisterAgent(ctx, &emptypb.Empty{})
//if err != nil {
//log.Println("could not register agent: ", err)
//os.Exit(1)
//}
//log.Println("agent registered with id: ", resp.AgentId)
