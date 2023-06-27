package main

import (
  "fmt"
  "log"
  "context"
  "time"
  "crypto/tls"

  "github.com/segmentio/kafka-go"
  "github.com/segmentio/kafka-go/sasl/plain"
)

func main() {
	 var target_records uint32
	 var counter uint32

	 target_records = 1000000

  mechanism := plain.Mechanism{
      Username: "IVB532DYA7CGGHTG",
      Password: "3j36pflBRfJtoQpQIH3hEtgAE7miFuirzMhwNLOrOfO98HL4Ls1gePg1mRDT1CpB",
  }

  dialer := &kafka.Dialer{
          TLS:       &tls.Config{
            MinVersion: tls.VersionTLS12,
          },
          SASLMechanism: mechanism,
        }

  r := kafka.NewReader(kafka.ReaderConfig{
    Brokers:   []string{"pkc-l7pr2.ap-south-1.aws.confluent.cloud:9092"},
    GroupID:   "consumer-group-id", // To read from all partitions of topic
    GroupTopics:   []string{"t3","t2","t1"},
    //Topic:     "perf-test-topic2",
    QueueCapacity: 10000,
    MinBytes:  100000, // 100KB
    MaxBytes:  10e6, // 10MB
    MaxWait:  time.Second,
    CommitInterval: 2*time.Second,
    StartOffset: kafka.FirstOffset,
    WatchPartitionChanges: true,
    Dialer:         dialer,
  })

  defer r.Close()

  fmt.Println("start consuming ... !!")

  startTime := time.Now()
  for {
	m, err := r.ReadMessage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	counter++

	fmt.Printf("counter:%v  message at topic:%v partition:%v offset:%v \n", counter, m.Topic, m.Partition, m.Offset)

	if counter  == target_records {
		fmt.Println("Target Records Read: ", target_records)
		break;
	}
	}

  endTime := time.Now()
  diff := endTime.Sub(startTime)
  fmt.Println("Time Taken in Seconds", diff.Seconds())
}
