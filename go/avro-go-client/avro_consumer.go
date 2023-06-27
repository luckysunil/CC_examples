// Reference: https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/confluent_cloud_example/confluent_cloud_example.go
package main

import (
  "fmt"
  "os"
  "time"
  "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
  "github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
  "github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/avro"
)

func main() {

  // Topic to receive data
  topic := "topic-go-avro"


  // ConfigMap is a map containing standard librdkafka configuration properties as documented in:
  // https://github.com/confluentinc/librdkafka/tree/master/CONFIGURATION.md

  consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
                      "bootstrap.servers":"pkc-l7pr2.ap-south-1.aws.confluent.cloud:9092",
                      "security.protocol":"SASL_SSL",
                      "sasl.mechanisms":"PLAIN",
                      "sasl.username":"544LWRBU2MBMP6RV",
                      "sasl.password":"nqIW1EeyskWH5Nkl0Yf1dNVtdlQqyc4AGYilU863bBFOAJY9Pi3Hmqxpqww9smKb",
                      "session.timeout.ms":6000,
                      "group.id":"my-group-avro",
                      "auto.offset.reset":"earliest",
                      })
  if err != nil {
    fmt.Printf("Failed to create consumer: %s\n", err)
    os.Exit(1)
  }
 fmt.Printf("Created Consumer %v\n", consumer)

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

  deser, err := avro.NewGenericDeserializer(srClient, serde.ValueSerde, avro.NewDeserializerConfig())
  if err != nil {
   fmt.Printf("Failed to create deserializer: %s\n", err)
   os.Exit(1)
 }

  topics := []string{topic}
  consumer.SubscribeTopics(topics, nil)

  for {
    message, err := consumer.ReadMessage(100 * time.Millisecond)
		if err == nil {
			received := User{}
			err := deser.DeserializeInto(*message.TopicPartition.Topic, message.Value, &received)

			if err != nil {
				fmt.Printf("Failed to deserialize payload: %s\n", err)
			} else {
				fmt.Printf("consumed from topic %s [%d] at offset %v: %+v \n",
					*message.TopicPartition.Topic,
					message.TopicPartition.Partition, message.TopicPartition.Offset,
					received)
			}
		}
	}

  consumer.Close()

}

// User is a simple record example
type User struct {
	Name           string `json:"name"`
	FavoriteNumber int64  `json:"favorite_number"`
	FavoriteColor  string `json:"favorite_color"`
}
