package kafka

import (
	"context"
	"github.com/arcosx/Sirius/config"
	"github.com/arcosx/Sirius/kafka/isolation"
	"github.com/arcosx/Sirius/util"
	"google.golang.org/grpc"
	"testing"
)

func init() {
	util.SetConfig()
	isolation.InitIsolationSet()

}

func init() {
	go NewKafkaService(config.C.Kafka.Rpc["port"])
}

// test function for
func Test_kafkaServer_ProduceAsync(t *testing.T) {
	ctx := context.Background()
	const address = "localhost:8848"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewKafkaClient(conn)
	t.Run("test ProduceAsync", func(t *testing.T) {
		c.ProduceAsync(ctx, &ProduceRequest{
			Isolation: "test",
			Topic:     "test",
			Message:   []byte("Test_kafkaServer_ProduceAsync"),
		})
	})
}
