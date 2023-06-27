// Reference: https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/confluent_cloud_example/confluent_cloud_example.go
package main

import (
  "fmt"
  "os"
  "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
  "github.com/linkedin/goavro/v2"
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
                      "group.id":"my-group-goavro",
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

  topics := []string{topic}
  consumer.SubscribeTopics(topics, nil)

  // https://pkg.go.dev/github.com/confluentinc/confluent-kafka-go/v2/schemaregistry#pkg-functions
  SchemaMetadata, err := srClient.GetLatestSchemaMetadata("topic-go-avro-value")
  if err != nil {
    fmt.Printf("Avro schema ID not found in message headers")
  }

  fmt.Printf("Schema: %s", SchemaMetadata.SchemaInfo.Schema)

  // https://pkg.go.dev/github.com/linkedin/goavro#NewCodec
  codec, err := goavro.NewCodec(SchemaMetadata.SchemaInfo.Schema)
  if err != nil {
    fmt.Printf("Codec creation failed")
  }

  for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}

    // https://pkg.go.dev/github.com/linkedin/goavro/v2#readme-usage
  	decodedMsgData, _, err := codec.NativeFromBinary(msg.Value[5:]) //First 4 byte is magic byte, so removing the same.
  	if err != nil {
  		fmt.Printf("Failed to decode Avro data: %s", err)
  	}

    fmt.Printf("Received message: topic=%s, partition=%d, offset=%d, Message: %+v\n",
        *msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset, decodedMsgData)

	}


  consumer.Close()

}
