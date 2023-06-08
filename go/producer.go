package main

import (
  "fmt"
  "context"
  "time"
  "crypto/tls"

  "github.com/segmentio/kafka-go"
  "github.com/segmentio/kafka-go/sasl/plain"
)

func main() {
  w := &kafka.Writer{
  	Addr:     kafka.TCP("pkc-l7pr2.ap-south-1.aws.confluent.cloud:9092"),
  	Topic:   "test-go",
    Transport: &kafka.Transport{
                TLS: &tls.Config{
                  MinVersion: tls.VersionTLS12,
                },
                SASL: plain.Mechanism{
                  Username: "AHZ2FDERNXU5MLKG",
                  Password: "szv19A1hDe3M1hIfEZcK+epZfGUR9UX9q+vbdJbPhatWAG9MddWGg5ivtA5xyyET",
                },
              },
  }
  defer w.Close()

  fmt.Println("start producing ... !!")
	for i := 0; i<3; i++ {
    fmt.Println(fmt.Sprint("Message: ", i))

		msg := kafka.Message{

			Value: []byte(fmt.Sprint("Message: ", i)),
		}
    fmt.Println("+")
		err := w.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println("failed to write messages:", err)
		} else {
			fmt.Println("produced", i) 
		}
		time.Sleep(1 * time.Second)
	}
}
