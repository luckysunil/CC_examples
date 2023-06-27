// Reference: https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/confluent_cloud_example/confluent_cloud_example.go
package main

import (
  "fmt"
  "os"
  "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/avro"
)

// User is a simple record example
type User struct {
	Name           string `json:"name"`
	FavoriteNumber int64  `json:"favorite_number"`
	FavoriteColor  string `json:"favorite_color"`
}

func main() {
  // ConfigMap is a map containing standard librdkafka configuration properties as documented in:
  // https://github.com/confluentinc/librdkafka/tree/master/CONFIGURATION.md

  producer, err := kafka.NewProducer(&kafka.ConfigMap{
                      "bootstrap.servers":"pkc-l7pr2.ap-south-1.aws.confluent.cloud:9092",
                      "security.protocol":"SASL_SSL",
                      "sasl.mechanisms":"PLAIN",
                      "sasl.username":"544LWRBU2MBMP6RV",
                      "sasl.password":"nqIW1EeyskWH5Nkl0Yf1dNVtdlQqyc4AGYilU863bBFOAJY9Pi3Hmqxpqww9smKb",
                      })
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
  // Close a Producer instance. The Producer object or its channels are no longer usable after this call.
  defer producer.Close()

  fmt.Printf("Created Producer %v\n", producer)

  // https://pkg.go.dev/github.com/confluentinc/confluent-kafka-go/v2/schemaregistry#Config
  srURL := "https://psrc-10dzz.ap-southeast-2.aws.confluent.cloud"
  srUsername := "434OPJN53I2C5ZZL"
  srPassword := "bkATLQ9koxOq5i7qzq83efZkkXjVNsHaThyKbVkQIeDf7SFo9bpyFaKGzv5MmEvC"

  //https://pkg.go.dev/github.com/confluentinc/confluent-kafka-go/v2/schemaregistry#NewConfigWithAuthentication
  srClient, err := schemaregistry.NewClient(schemaregistry.NewConfigWithAuthentication(srURL, srUsername, srPassword))
  if err != nil {
    fmt.Printf("Failed to create schema registry client: %s\n", err)
    os.Exit(1)
  }

 // https://pkg.go.dev/github.com/olli-ai/confluent-kafka-go/schemaregistry/serde/avro#NewGenericSerializer
  ser, err := avro.NewGenericSerializer(srClient, serde.ValueSerde, avro.NewSerializerConfig())
	if err != nil {
		fmt.Printf("Failed to create serializer: %s\n", err)
		os.Exit(1)
	}

  // Topic to send data
  topic := "topic-go-avro"

  // Sample data to be avro serialized and send to topic
  value := User {
    Name:           "Second user",
    FavoriteNumber: 1,
    FavoriteColor:  "Green",
 }

 fmt.Println("start producing ... !!")

 payload, err := ser.Serialize(topic, &value)
	if err != nil {
		fmt.Printf("Failed to serialize payload: %s\n", err)
		os.Exit(1)
	}

  // https://pkg.go.dev/gopkg.in/confluentinc/confluent-kafka-go.v1@v1.8.2/kafka#hdr-Producer_events
  producer.Produce(&kafka.Message{
  		  TopicPartition: kafka.TopicPartition{Topic: &topic,
  			Partition: kafka.PartitionAny},
  		  Value: payload}, nil)

  	// Wait for delivery report
  	e := <-producer.Events()

  	message := e.(*kafka.Message)
  	if message.TopicPartition.Error != nil {
  		fmt.Printf("failed to deliver message: %v\n",
  			message.TopicPartition)
  	} else {
  		fmt.Printf("delivered to topic %s [%d] at offset %v\n",
  			*message.TopicPartition.Topic,
  			message.TopicPartition.Partition,
  			message.TopicPartition.Offset)
  	}

}
