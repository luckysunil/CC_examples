#!/bin/bash
exec 1>perf_test_result.txt 2>&1

echo "This is a Kafka Performance Test script for a single Client VM"
echo "=============================================================="

# counter for counting the number of tests performed
counter=0

# Please enter the test runtime in mins:(e.g. 15, 60)
runtime=1

echo "Benchmark runtime for each case in mins:" $runtime

# Enter Message size in bytes
msgbytes=(41984)
echo ${msgbytes[*]}

# Enter Throughput in  mb/s: (e.g. 5, 15, 50, 100)
tp_value=(10)
echo ${tp_value[*]}

echo "Producer Properties:"
# acks: 0,1,all
acks_value=(all)
echo ${acks_value[*]}

# batch.size: 100000–200000
batch_size=16384

# linger.ms: 10-100
linger_ms=0

# compression.type: lz4, snappy, zstd, gzip
compression_type=none

echo "Kafka Properties:"
bootstrap_server=pkc-7prvp.centralindia.azure.confluent.cloud:9092
topic_name=perf-test

echo "Topic Properties"
partitions=(1)
echo ${partitions[*]}

for a in ${acks_value[*]}
do
  for p in ${partitions[*]}
  do
    for m in ${msgbytes[*]}
    do
      for t in ${tp_value[*]}
      do
          #Throuhput calculation
          tp=$(($t*1048576/$m))
          #echo "throughput param:" $tp

          #Benchmark test runtime duration
          test_records=$(($tp*$runtime*60))
          echo "test_records:" $test_records

          # Start Time for the TEST
          $((++counter))
          start_date_time="`date "+%Y-%m-%d %H:%M:%S"`";
          echo "TEST:"$counter "starttime:" $start_date_time "Producer properties# acks:" $a "linger.ms:" $linger_ms "compression.type:"$compression_type  "partitions:" $p "record-size:" $m "throughput" $t


          ./kafka-topics --bootstrap-server $bootstrap_server  --command-config client.config --topic $topic_name --delete

          sleep 1

          ./kafka-topics  --bootstrap-server $bootstrap_server   --command-config client.config --create --topic $topic_name --partitions $p --replication-factor 3

  	      sleep 1

          ./kafka-producer-perf-test --topic $topic_name --record-size $m --producer.config client.config --throughput $tp --num-records $test_records --producer-props acks=$a linger.ms=$linger_ms compression.type=$compression_type | tee /tmp/producer &

          ./kafka-consumer-perf-test --broker-list $bootstrap_server --consumer.config client.config --topic $topic_name --messages $test_records --group=perf-test-1 --show-detailed-stats --hide-header --timeout 60000 | tee /tmp/consumer1 &

          ./kafka-consumer-perf-test --broker-list $bootstrap_server --consumer.config client.config --topic $topic_name --messages $test_records --group=perf-test-2 --show-detailed-stats --hide-header --timeout 60000 | tee /tmp/consumer2 &

          ./kafka-consumer-perf-test --broker-list $bootstrap_server --consumer.config client.config --topic $topic_name --messages $test_records --group=perf-test-3 --show-detailed-stats --hide-header --timeout 60000 | tee /tmp/consumer3 &

          wait

	        echo "TEST:"$counter "Result:"$(tail -n 1 /tmp/producer | sed 's|.(\(.\))|Producer: \1|g')

          echo "TEST:"$counter "Consumer 1:" `cat /tmp/consumer1 | awk -F"," '{if($8>0){msec+=$8};mb=$3}END{print mb*1000/msec}'` "MB/sec"

          echo "TEST:"$counter "Consumer 2:" `cat /tmp/consumer2 | awk -F"," '{if($8>0){msec+=$8};mb=$3}END{print mb*1000/msec}'` "MB/sec"

          echo "TEST:"$counter "Consumer 3:" `cat /tmp/consumer3 | awk -F"," '{if($8>0){msec+=$8};mb=$3}END{print mb*1000/msec}'` "MB/sec"

          # End Time for the TEST
          end_date_time="`date "+%Y-%m-%d %H:%M:%S"`";
          echo "TEST:"$counter "endtime:" $end_date_time

          wait

          sleep 1
      done
    done
  done
done
