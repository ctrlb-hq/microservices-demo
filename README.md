## Intro
This is a sample microservices app just like sock-shop and online boutique.

## Install
`kubectl apply -f K8/manifest.yaml`
`kubectl port-forward -n atlas service/python-frontend-service 30000 --address=0.0.0.0`

## Services

### python-frontend-service
1. Endpoint: python-service:30000?/?number=10.
    a. Puts the request on kafka queue with this number.
    b. Returns a uuid to fetch result.
2. Endpoint: python-service:30000/getResult?uuid=3f641773-4d54-44ee-b2eb-290b215116b7.
    a. Calls java-service:30001/numbers/{uid} and asks it to get the result.
    b. Returns square of the input number multiplied by 2.

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
