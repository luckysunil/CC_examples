#!/bin/bash

<<comment
  This script can be used to generate multiple topics with specific number of
  records per topic per partition to emulate a deployment environment
comment

total_topics=3
partitions_per_topic=3
msgs_per_topic=100000

bootstrap_server=pkc-l7pr2.ap-south-1.aws.confluent.cloud:9092
msgbytes=5120

throughput=-1 #No throttling

# acks: 0,1,all
acks_value=all

# batch.size: 100000–200000
batch_size=16384

# linger.ms: 10-100
linger_ms=0

# compression.type: lz4, snappy, zstd, gzip
compression_type=none

echo "Total no. of topics: $total_topics"

for ((i=total_topics; i>=1; i--))
do
  topic_name="t$i"
  echo "$topic_name"
  echo "*********************"

  topic_list+=$topic_name","

  ./kafka-topics --bootstrap-server $bootstrap_server  --command-config client.config --topic $topic_name --delete

  sleep 1

  ./kafka-topics  --bootstrap-server $bootstrap_server   --command-config client.config --create --topic $topic_name --partitions $partitions --replication-factor 3

  sleep 1

  ./kafka-producer-perf-test --topic $topic_name --record-size $msgbytes --producer.config client.config --throughput $throughput --num-records $msgs_per_topic --producer-props acks=$acks_value linger.ms=$linger_ms compression.type=$compression_type

done

echo "Topic list: [ $topic_list ]"
