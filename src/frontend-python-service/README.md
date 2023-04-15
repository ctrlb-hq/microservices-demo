### Steps.
1. docker build -f Dockerfile -t python-frontend-service .
2. docker tag python-frontend-service dev0zklabs/atlas-demo-microservice:python-frontend-service
3. docker push dev0zklabs/atlas-demo-microservice:python-frontend-service
4. docker run -p 30000:30000 --network host --add-host 0.0.0.0:0.0.0.0 python-frontend-service