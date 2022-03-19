
const { Kafka } = require('kafkajs')

// Creating client instance to connect to the Confluent Cloud
const kafka = new Kafka({
  clientId: 'demo-producer',
  brokers: ['pkc-7wxyz.centralindia.azure.confluent.cloud:9092'],
  ssl: true,
  sasl: {
    mechanism: 'plain',
    username: 'CY7L72ZT2YW3MJ3I',
    password: 'Ph2yBwaRumvBx+rggl8EM489WJx1V1Lm47Tp6LOD0AV0VXAUPHQWjf3YNWKgfwoV'
  }
})

const producer = kafka.producer()

producer.on('producer.connect', () => {
  console.log('KafkaProvider: connected');
});

producer.on('producer.disconnect', () => {
  console.log('KafkaProvider: disconnected');
});

producer.on('producer.network.request_timeout', (payload) => {
  console.log(`KafkaProvider: request timeout ${payload.clientId}`);
});

const run = async () => {
  await producer.connect()
  await producer.send({
    topic: 'test-topic', //Replace with appropriate topic
    messages: [
      {
        value: Buffer.from(JSON.stringify(
          {
            "student_name": "Sunil",
            "user_id": 1234,
            "performance": {
              "assessment": {
                "institute": "ABC",
                "subject": "Physics",
                "grade": 10.0,
              }
            }
          }
        ))
      },
    ],
    acks:-1, //-1 = all insync replicas must acknowledge. Details: https://kafka.js.org/docs/producing
  })
}

run().catch(console.error)
