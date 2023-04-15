## Steps
https://levelup.gitconnected.com/how-to-deploy-apache-kafka-with-kubernetes-9bd5caf7694f 

1. Create namespace: `kubectl apply -f 00-namespace.yaml`
2. Deploy zookeeper: `kubectl apply -f 01-zookeeper.yaml`
3. Deploy Kafka Broker: `kubectl apply -f 02-kafka.yaml`
4. Notice the line in 02-kafka.yaml where we provide a value for KAFKA_ADVERTISED_LISTENERS. To ensure that Zookeeper and Kafka can communicate by using this hostname (kafka-broker), we need to add the following entry to the /etc/hosts file on our local machine:
`127.0.0.1 kafka-broker`


## Testing Kafka Topics
1. kubectl port-forward -n kafka service/kafka-service 9092
2. sudo apt-get install kafkacat
3. echo "hello world!" | kafkacat -P -b localhost:9092 -t test
4. kafkacat -C -b localhost:9092 -t test