Make sure to have Java installed
$ java -version
openjdk version "11.0.11" 2021-04-20
OpenJDK Runtime Environment (build 11.0.11+9-Ubuntu-0ubuntu2.18.04)
OpenJDK 64-Bit Server VM (build 11.0.11+9-Ubuntu-0ubuntu2.18.04, mixed mode, sharing)

You can install using below command:
sudo apt install default-jre

Go to Confluent Platform Download link below:
https://docs.confluent.io/platform/current/installation/installing_cp/zip-tar.html

Download the Latest version using command like below:
curl -O http://packages.confluent.io/archive/7.0/confluent-7.0.2.tar.gz

Untar the file
tar xzf confluent-6.2.0.tar.gz

Create a client.config file in the confluent-7.0.2/bin folder, and add the below content:

bootstrap.servers=pkc-192zz.asia-south1.gcp.confluent.cloud:9092
security.protocol=SASL_SSL
sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required username="UFPQT5ARLENUMMTP" password="dQjFLUQ+h/1W8a3UGue4pNTJIYN6LMvM7VcPviyPJuqjIIC+8Gr3aFRcNXeb4GIF";
sasl.mechanism=PLAIN

Copy the cc-perf-test-script.sh file into the confluent-7.0.2/bin folder and modify the parameters accordingly.

Make this file executable:
chmod +x cc-perf-test-script.sh
