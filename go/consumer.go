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

  mechanism := plain.Mechanism{
      Username: "AHZ2FDERNXU5MLKG",
      Password: "szv19A1hDe3M1hIfEZcK+epZfGUR9UX9q+vbdJbPhatWAG9MddWGg5ivtA5xyyET",
  }

  dialer := &kafka.Dialer{
          TLS:       &tls.Config{
            MinVersion: tls.VersionTLS12,
          },
          SASLMechanism: mechanism,
        }

  r := kafka.NewReader(kafka.ReaderConfig{
    Brokers:   []string{"pkc-l7pr2.ap-south-1.aws.confluent.cloud:9092"},
    Topic:     "test-go",
    QueueCapacity: 10000,
    MinBytes:  1, // 10KB
    MaxBytes:  10e6, // 10MB
    CommitInterval: time.Second,
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
    if m.Offset == -1 {
      fmt.Println("Something incorrect here")
      break;
    }
		//fmt.Printf("message at topic:%v partition:%v offset:%v	msg: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))

    lag, err := r.ReadLag(context.Background())
    //fmt.Println("lag: %d", lag)

    if (lag == 0) {
      fmt.Println("All records read")
      break;
    }
	}
  endTime := time.Now()
  diff := endTime.Sub(startTime)
  fmt.Println(diff.Seconds())
}
