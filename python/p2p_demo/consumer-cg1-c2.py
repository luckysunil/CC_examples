# =============================================================================
#
# Consume messages from Confluent Cloud
# Using Confluent Python Client for Apache Kafka
#
# This is a quick demo derived from original repository at
# https://github.com/confluentinc/confluent-kafka-python
# =============================================================================

#from confluent_kafka import Consumer
from confluent_kafka import DeserializingConsumer
from confluent_kafka.serialization import StringDeserializer
import json

if __name__ == '__main__':


    topic = 'demo-topic'  # Add the topic created on Confluent Cloud

    # Create Consumer instance
    # 'auto.offset.reset=earliest' to start reading from the beginning of the

    config_data = {
                    'client.id':'consumer-cg1-c2', #Client Identifier

                    'bootstrap.servers':'pkc-41p56.asia-south1.gcp.confluent.cloud:9092', # Add the Bootstrap server info from Confluent Cloud
                    'security.protocol':'SASL_SSL',
                    'sasl.mechanisms':'PLAIN',
                    'sasl.username':'<ADD CLUSTER API KEY>', # Add Cluster API Key
                    'sasl.password':'<ADD CLUSTER API SECRET>', # Add Cluster API Secret

                    'group.id':'cg1',
                    'auto.offset.reset':'earliest',
                    'key.deserializer': StringDeserializer('utf_8'),
                    'value.deserializer': StringDeserializer('utf_8')
                  }

    consumer = DeserializingConsumer(config_data)

    # Subscribe to topic
    consumer.subscribe([topic])

    # Process messages
    total_count = 0
    try:
        while True:
            msg = consumer.poll(1.0)

            if msg is None:
                # No message available within timeout.
                # Initial message consumption may take up to
                # `session.timeout.ms` for the consumer group to
                # rebalance and start consuming

                #print("Waiting for message.")
                continue
            elif msg.error():
                print('error: {}'.format(msg.error()))
            else:
                # Check for Kafka message
                key = msg.key()


                value = msg.value()
                #value = json.loads(record_value)
                #count = data['count']
                #total_count += count
                print("Consumed record:\nPartition:{}\nOffset:{}\nKey:{}\nValue:{}\n"
                      .format(msg.partition(), msg.offset(), key, value))
    except KeyboardInterrupt:
        pass
    finally:
        # Leave group and commit final offsets
        consumer.close()
