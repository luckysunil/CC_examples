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
# Consume messages from Confluent Cloud
# Using Confluent Python Client for Apache Kafka
# Reads Avro data, integration with Confluent Cloud Schema Registry
#
# =============================================================================

from confluent_kafka import DeserializingConsumer
from confluent_kafka.schema_registry import SchemaRegistryClient
from confluent_kafka.schema_registry.avro import AvroDeserializer
from confluent_kafka.serialization import StringDeserializer
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
    user_avro_deserializer = AvroDeserializer(schema_registry_client,
                                            clients_common.schema_str,
                                            clients_common.User.dict_to_user)


    config_data = {
                'client.id':'consumer-c1', #Client Identifier

                'bootstrap.servers':'<BOOTSTRAP SERVER URL>', # Add the Bootstrap server info from Confluent Cloud
                'security.protocol':'SASL_SSL',
                'sasl.mechanisms':'PLAIN',
                'sasl.username':'<ADD CLUSTER API KEY>', # Add Cluster API Key
                'sasl.password':'<ADD CLUSTER API SECRET>', # Add Cluster API Secret

                'key.deserializer': StringDeserializer('utf_8'),
                'value.deserializer': user_avro_deserializer,

                'group.id':'cg',
                'auto.offset.reset':'earliest'
              }

    consumer = DeserializingConsumer(config_data)

    consumer.subscribe([topic])

    while True:
        try:
            # SIGINT can't be handled when polling, limit timeout to 1 second.
            msg = consumer.poll(1.0)
            if msg is None:
                continue

            user = msg.value()
            if user is not None:
                print("User record:\n"
                      "\tname: {}\n"
                      "\tfavorite_number: {}\n"
                      "\tfavorite_color: {}\n"
                      .format(user.name,
                              user.favorite_color,
                              user.favorite_number))
        except KeyboardInterrupt:
            break

    consumer.close()
