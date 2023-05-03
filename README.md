## Intro
This is a sample microservices app just like sock-shop and online boutique.

## Install
### Pre-reqs
Install opentelemetry instrumentation
```
https://www.notion.so/dev0ai/Tutorial-Otel-Autoinstrumentation-4a0e446b22b04f6291693c20fd49092b 
```
Create go instrumentation docker image
```
https://www.notion.so/dev0ai/Tutorial-Instrumenting-go-applications-b9793c39276946e9978ec2087f9d670a?pvs=4#a2a093bb9cfb45398e40d965c7dc2ced
```
### Main install
`kubectl apply -f K8/manifest.yaml`
`kubectl port-forward -n atlas service/python-frontend-service 30000 --address=0.0.0.0`

## Services

### python-frontend-service
1. Endpoint: python-service:30000/?number=10.  
    a. Puts the request on kafka queue with this number.  
    b. Returns a uuid to fetch result.  
2. Endpoint: python-service:30000/result
    a. form body must have uuid=3f641773-4d54-44ee-b2eb-290b215116b7.  
    b. Calls java-service:30001/numbers/{uid} and asks it to get the result.  
    c. Returns square of the input number multiplied by 2.  
3. Endpoint: python-service:30000/helloGo  
    a. Calls go-service:30002/ping/  

### kafka
1. A kafka queue where python service puts the request and java service fetches this request.

### db-service
1. java service puts the square of number obtained from kafka on db. 
2. go-service gets this number when asked by java-service to obtain result.

### java-service
1. Kafka listener which keeps consuming from kafka queue. It squares this number and puts on db.  
2. Endpoint: localhost:30000/getResult?uuid=3f641773-4d54-44ee-b2eb-290b215116b7  
    a. Calls go-service:30002/fetchNumber?uuid=3f641773-4d54-44ee-b2eb-290b215116b7 and asks for result  
    b. Returns the number recived by go-service after multiplying it by 2.  

### go-service
1. Endpoint: go-service:30002/fetchNumber?uuid=3f641773-4d54-44ee-b2eb-290b215116b7  
    a. Fetched this uuid from DB.  
    b. Multiplies it by 2 and returns.  
2. Endpoint: go-service:30002/ping  
    a. Simply returns pong  
3. Endpoint: go-service:30002/pingPython  
    a. Calls python-service:30000/helloGo  

### node-service
TODO - Add documentation Adarsh
