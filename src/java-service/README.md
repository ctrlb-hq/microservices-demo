## Steps
1. Export parameters:
export PYTHON_SERVICE_HOST=0.0.0.0
export PYTHON_SERVICE_PORT="30000"
export JAVA_SERVICE_HOST=0.0.0.0
export JAVA_SERVICE_PORT="30001"
export GO_SERVICE_HOST=0.0.0.0
export GO_SERVICE_PORT="30002"
export DB_SERVICE_HOST=0.0.0.0
export DB_SERVICE_PORT="5042"
export KAFKA_SERVICE_HOST=0.0.0.0
export KAFKA_SERVICE_PORT="9092"
export KAFKA_SERVICE_TOPIC=test
export POSTGRES_DB=postgresdb
export POSTGRES_USER=admin
export POSTGRES_PASSWORD=psltest
export DB_SERVICE_NAME=db-service

2. Install jdk20. sudo apt install openjdk-17-jdk

3. ./mvnw spring-boot:run