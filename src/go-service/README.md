### Steps.
1. docker build -f Dockerfile -t go-service .
2. docker tag go-service dev0zklabs/atlas-demo-microservice:go-service
3. docker push dev0zklabs/atlas-demo-microservice:go-service
4. docker run --network host --add-host 0.0.0.0:0.0.0.0 go-service