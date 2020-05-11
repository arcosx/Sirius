package kafka

import (
	"context"
	"github.com/arcosx/Sirius/config"
	"github.com/arcosx/Sirius/kafka/isolation"
	"github.com/arcosx/Sirius/util"
	"google.golang.org/grpc"
	"strconv"
	"testing"
	"time"
)

func init() {
	util.SetConfig()
	isolation.InitIsolationSet()

}

// init a real grpc for test
func initService() {
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

func Test_kafkaServer_ProduceStream(t *testing.T) {
	ctx := context.Background()
	const address = "localhost:8848"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewKafkaClient(conn)
	t.Run("test ProduceStream", func(t *testing.T) {
		kafkaProduceStreamClient, err := c.ProduceStream(ctx)
		if err != nil {
			t.Error(err)
		}
		sendTimes := 0
		for {
			time.Sleep(2 * time.Second)
			sendTimes += 1
			kafkaProduceStreamClient.Send(&ProduceRequest{
				Isolation: "test",
				Topic:     "test",
				Message:   []byte("Test_kafkaServer_ProduceAsync" + " times:" + strconv.Itoa(sendTimes) + util.GetTimeNowString()),
			})
			if sendTimes == 10 {
				kafkaProduceStreamClient.CloseSend()
				break
			}
		}
	})
}

func Test_kafkaServer_ConsumeStream(t *testing.T) {
	initService()
	ctx := context.Background()
	const address = "localhost:8848"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewKafkaClient(conn)
	t.Run("test ConsumeStream", func(t *testing.T) {
		kafkaConsumer, err := c.ConsumeStream(ctx, &ConsumeRequest{
			Isolation: "test",
			Topic:     "test",
		})
		if err != nil {
			t.Error(err)
		}
		for {
			consumeRes, _ := kafkaConsumer.Recv()
			t.Log(consumeRes)
		}
	})
}
