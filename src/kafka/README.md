## Steps
1. docker-compose up
2. Find the container id using docker container ls
3. Create topic: docker exec -it 1217a863f2ff /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --topic first_topic --create --partitions 3 --replication-factor 1