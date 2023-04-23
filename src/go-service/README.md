### Steps.
1. docker build -f Dockerfile -t go-service .
2. docker tag go-service dev0zklabs/atlas-demo-microservice:go-service
3. docker push dev0zklabs/atlas-demo-microservice:go-service
4. docker run --network host --add-host 0.0.0.0:0.0.0.0 go-service

### Instrumentation
opentelemetry auto-instrumentation is wip at the moment.
Checkout https://www.notion.so/dev0ai/Tutorial-Instrumenting-go-applications-b9793c39276946e9978ec2087f9d670a to learn more