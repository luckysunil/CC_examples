# =============================================================================
#
# Produces messages to Confluent Cloud
# Using Confluent Python Client for Apache Kafka
#
# This is a quick demo derived from original repository at
# https://github.com/confluentinc/confluent-kafka-python
# =============================================================================

#from confluent_kafka import Producer, KafkaError
from confluent_kafka import SerializingProducer, KafkaError
from confluent_kafka.serialization import StringSerializer
import json

if __name__ == '__main__':

    kafka_topic = 'demo-topic'  # Add the topic created on Confluent Cloud

    config_data = {
                    'client.id':'producer-p1', #Client Identifier

                    'bootstrap.servers':'pkc-41p56.asia-south1.gcp.confluent.cloud:9092', # Add the Bootstrap server info from Confluent Cloud
                    'security.protocol':'SASL_SSL',
                    'sasl.mechanisms':'PLAIN',
                    'sasl.username':'<ADD CLUSTER API KEY>', # Add Cluster API Key
                    'sasl.password':'<ADD CLUSTER API SECRET>', # Add Cluster API Secret

                    'key.serializer': StringSerializer('utf_8'),
                    'value.serializer': StringSerializer('utf_8')
                  }

    producer = SerializingProducer(config_data)

    delivered_records = 0

    def acked(err, msg):
        global delivered_records
        """Delivery report handler called on
        successful or failed delivery of message
        """
        if err is not None:
            print("Failed to deliver message: {}".format(err))
        else:
            delivered_records += 1
            print("Produced record to topic {} partition [{}] @ offset {} \n"
                  .format(msg.topic(), msg.partition(), msg.offset()))

    while(True):

        print("\n\nCreate record data {\"Key\":\"Value\"} ")
        key = input("\nEnter Key: ")
        value = input("Enter Value: ")

        record_key = key
        record_value = value

        print("\n Producing record ==> \"{}\"\t{} \n".format(record_key, record_value))
        producer.produce(kafka_topic, key=record_key, value=record_value, on_delivery=acked)
        producer.flush()

    print("{} messages were produced to topic {}!".format(delivered_records, kafka_topic))
