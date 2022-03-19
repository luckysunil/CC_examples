const { Kafka } = require('kafkajs')

// Creating client instance to connect to the Confluent Cloud
const kafka = new Kafka({
  brokers: ['pkc-7wxyz.centralindia.azure.confluent.cloud:9092'],
  ssl: true,
  sasl: {
    mechanism: 'plain',
    username: 'CY7L72ZT2YW3MJ3I',
    password: 'Ph2yBwaRumvBx+rggl8EM489WJx1V1Lm47Tp6LOD0AV0VXAUPHQWjf3YNWKgfwoV'
  }
})

const consumer = kafka.consumer({groupId: 'demo-consumer-groupId-2'})

consumer.on('consumer.connect', () => {
  console.log('KafkaProvider: connected');
});

const run = async () => {
	// first, we wait for the client to connect and subscribe to the given topic
	await consumer.connect()

  //Subscribe to appropriate topic
	await consumer.subscribe({ topic:'test-topic', fromBeginning: true })

  await consumer.run({
    autoCommit: false,
    eachMessage: async ({ topic, partition, message }) => {
      console.log({
        topic,
        partition,
        offset: message.offset,
        value: message.value.toString(),
      })

      await consumer.commitOffsets([{ topic, partition, offset: (Number(message.offset) + 1).toString() }]);
    },
  })
}

run().catch(console.error)
