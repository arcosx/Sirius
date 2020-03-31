// Direct implementation of the gRPC method
package kafka

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type kafkaServer struct {
	UnimplementedKafkaServer
}

func NewKafkaServer(addr string, opts ...grpc.ServerOption) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Error("Create gRPC failed", err)
		return err
	}
	grpcServer := grpc.NewServer(opts...)
	RegisterKafkaServer(grpcServer, &kafkaServer{})
	grpcServer.Serve(lis)
	return nil
}

func (k *kafkaServer) Produce(req *ProduceRequest, stream Kafka_ProduceServer) error {
	return nil
}

func (k *kafkaServer) Consume(req *ConsumeRequest, stream Kafka_ConsumeServer) error {
	return nil
}
