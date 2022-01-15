#!/usr/bin/env python
#
# Copyright 2020 Confluent Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# =============================================================================
#
# Produce messages to Confluent Cloud
# Using Confluent Python Client for Apache Kafka
# Writes Avro data, integration with Confluent Cloud Schema Registry
#
# =============================================================================
from confluent_kafka import SerializingProducer
from confluent_kafka.schema_registry import SchemaRegistryClient
from confluent_kafka.serialization import StringSerializer
from confluent_kafka.schema_registry.avro import AvroSerializer
from uuid import uuid4

import clients_common



if __name__ == '__main__':

    # Read arguments and configurations and initialize

    topic = 'demo-topic' #Create topic in Confluent Cloud

    schema_registry_conf = {
        'url': '<SCHEMA REGISTRY URL>',
        'basic.auth.user.info': '<SR_API_KEY:SR_API_SECRET>' # basic.auth.user.info={{ SR_API_KEY }}:{{ SR_API_SECRET }}
        }

    schema_registry_client = SchemaRegistryClient(schema_registry_conf)
    user_avro_serializer = AvroSerializer(schema_registry_client,
                                        clients_common.schema_str,
                                        clients_common.User.user_to_dict)


    config_data = {
                'client.id':'producer-p1', #Client Identifier

                'bootstrap.servers':'<BOOTSTRAP SERVER URL>', # Add the Bootstrap server info from Confluent Cloud
                'security.protocol':'SASL_SSL',
                'sasl.mechanisms':'PLAIN',
                'sasl.username':'<ADD CLUSTER API KEY>', # Add Cluster API Key
                'sasl.password':'<ADD CLUSTER API SECRET>', # Add Cluster API Secret

                'key.serializer': StringSerializer('utf_8'),
                'value.serializer': user_avro_serializer
              }

    producer = SerializingProducer(config_data)

    delivered_records = 0

    # Optional per-message on_delivery handler (triggered by poll() or flush())
    # when a message has been successfully delivered or
    # permanently failed delivery (after retries).
    def acked(err, msg):
        global delivered_records
        """Delivery report handler called on
        successful or failed delivery of message
        """
        if err is not None:
            print("Failed to deliver message: {}".format(err))
        else:
            delivered_records += 1
            print("Produced record to topic {} partition [{}] @ offset {}"
                  .format(msg.topic(), msg.partition(), msg.offset()))

    for n in range(1):
        user_name = input("Enter name: ")
        user_favorite_number = int(input("Enter favorite number: "))
        user_favorite_color = input("Enter favorite color: ")
        user = clients_common.User(name=user_name,
                    favorite_color=user_favorite_color,
                    favorite_number=user_favorite_number)

        #producer.produce(topic=topic, key=str(uuid4()), value=count_object, on_delivery=acked)
        producer.produce(topic=topic, value=user, on_delivery=acked)
        producer.poll(0)

    producer.flush()

    print("{} messages were produced to topic {}!".format(delivered_records, topic))
