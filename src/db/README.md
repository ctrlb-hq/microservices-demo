## Steps
1. docker build -f Dockerfile -t db-service .
2. docker tag db-service dev0zklabs/atlas-demo-microservice:db-service
3. docker push dev0zklabs/atlas-demo-microservice:db-service

## How to run
1. docker run --network host --add-host 0.0.0.0:0.0.0.0 db-service
