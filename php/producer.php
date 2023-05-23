<?php

        echo "Testing PHP Client for Confluent Cloud \n";

        $conf = new RdKafka\Conf();
        $conf->set('sasl.mechanisms', 'PLAIN');
        $conf->set('security.protocol', 'SASL_SSL');
        $conf->set('sasl.username', 'NWB7RZ7A62X4MLSQ');
        $conf->set('sasl.password', 'G2mkxaZowELPZATreuOKDPX6i8Y/pLkDML3TXDtr6VlKJC0EzeV6pkG4oReVFjTg');
        $conf->set('bootstrap.servers', 'pkc-7prvp.centralindia.azure.confluent.cloud:9092');

        $producer = new RdKafka\Producer($conf);

        $topic = $producer->newTopic("test-topic");

        for ($i = 0; $i < 10; $i++) {
                $topic->produce(RD_KAFKA_PARTITION_UA, 0, "Message $i");
                echo "Message $i \n";
                $producer->poll(0);
        }

        for ($flushRetries = 0; $flushRetries < 10; $flushRetries++) {
                        $result = $producer->flush(10000);
                if (RD_KAFKA_RESP_ERR_NO_ERROR === $result) {
                        break;
                }
        }

        if (RD_KAFKA_RESP_ERR_NO_ERROR !== $result) {
                throw new \RuntimeException('Was unable to flush, messages might be lost!');
        }

?>
