#!/bin/bash

bootstrap_server=pkc-l7pr2.ap-south-1.aws.confluent.cloud:9092
topic_name=perf-test-topic
partitions=1
msgbytes=5120
num_records=1000000 #Total records to be produced
throughput=-1 #No throttling

# acks: 0,1,all
acks_value=all

# batch.size: 100000–200000
batch_size=16384

# linger.ms: 10-100
linger_ms=0

# compression.type: lz4, snappy, zstd, gzip
compression_type=none


./kafka-topics --bootstrap-server $bootstrap_server  --command-config client.config --topic $topic_name --delete

sleep 1

./kafka-topics  --bootstrap-server $bootstrap_server   --command-config client.config --create --topic $topic_name --partitions $partitions --replication-factor 3

sleep 1

./kafka-producer-perf-test --topic $topic_name --record-size $msgbytes --producer.config client.config --throughput $throughput --num-records $num_records --producer-props acks=$acks_value linger.ms=$linger_ms compression.type=$compression_type
