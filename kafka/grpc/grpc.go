// Direct implementation of the gRPC method
package kafka

import (
	"context"
	"github.com/arcosx/Sirius/kafka/isolation"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type kafkaServer struct {
	UnimplementedKafkaServer
}

func NewKafkaService(addr string, opts ...grpc.ServerOption) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Create Kafka Service gRPC failed:Failed net Listen", err)
	}
	grpcServer := grpc.NewServer(opts...)
	RegisterKafkaServer(grpcServer, &kafkaServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Create Kafka Service gRPC failed:Failed to serve: %v", err)
	}
	log.Info("Create Kafka Service gRPC success on", addr)
}

func (k *kafkaServer) ProduceAsync(ctx context.Context, request *ProduceRequest) (*Empty, error) {
	isolationKeyWord := request.Isolation
	isolationInstance, err := isolation.IsolationSet.GetIsolation(isolationKeyWord)
	if err != nil {
		log.Error("ProduceAsync GetIsolation", isolationKeyWord, "Error", err)
		return nil, err
	}
	isolationInstance.ProduceAsync(request.Topic, request.Message)
	return new(Empty), nil
}
