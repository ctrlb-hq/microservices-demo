# Importing flask module in the project is mandatory
# An object of Flask class is our WSGI application.
from flask import Flask, request
import os
import uuid
import requests

PYTHON_SERVICE_HOST = os.getenv("PYTHON_SERVICE_HOST","0.0.0.0")
PYTHON_SERVICE_PORT = os.getenv("PYTHON_SERVICE_PORT","30000")


JAVA_SERVICE_HOST = os.getenv("JAVA_SERVICE_HOST","0.0.0.0")
JAVA_SERVICE_PORT = os.getenv("JAVA_SERVICE_PORT","30001")

GO_SERVICE_HOST = os.getenv("GO_SERVICE_HOST","0.0.0.0")
GO_SERVICE_PORT = os.getenv("GO_SERVICE_PORT","30002")

KAFKA_SERVICE_HOST = os.getenv("KAFKA_SERVICE_HOST","0.0.0.0")
KAFKA_SERVICE_PORT = os.getenv("KAFKA_SERVICE_PORT","9092")
KAFKA_SERVICE_TOPIC = os.getenv("KAFKA_SERVICE_TOPIC","test")

app = Flask(__name__)

from kafka import KafkaProducer
producer = KafkaProducer(
    bootstrap_servers=[f"{KAFKA_SERVICE_HOST}:{KAFKA_SERVICE_PORT}"],
    value_serializer=lambda x: str(x).encode('utf-8')
)

"""
Frontend calls this function with a get parameter 'number'
This function puts this on kafka and gives back a uuid for the result 
"""
@app.route('/')
def home():
    number = request.args.get('number')
    uid = str(uuid.uuid4())
    data = {'uuid': uid, 'number':number}
    producer.send(KAFKA_SERVICE_TOPIC, value=data)
    return str(f"Your request has been registered against {uid}")

"""
This function calls the Java service endpoint /numbers with uuid parameter as obtained here
Then prints the result.
"""
@app.route('/getResult')
def result():
    uid = request.args.get('uuid')
    PARAMS = {'uuid':uid}

    # sending get request and saving the response as response object
    r = requests.get(url = f"http://{JAVA_SERVICE_HOST}:{JAVA_SERVICE_PORT}/numbers/{uid}", params = PARAMS)
    # extracting data in json format
    data = r.json()
    return str(f"Your response is: {data}")

"""
This function calls the Java service endpoint /numbers with uuid parameter as obtained here
Then prints the result.
"""
@app.route('/helloGo')
def helloGo():
    # sending get request and saving the response as response object
    r = requests.get(url = f"http://{GO_SERVICE_HOST}:{GO_SERVICE_PORT}/ping/")
    # extracting data in json format
    print(r)
    return str(f"Your response is: {r.text}")

"""
This function calls the Java service endpoint /numbers with uuid parameter as obtained here
Then prints the result.
"""
@app.route('/ping')
def ping():
    # sending get request and saving the response as response object
    r = requests.get(url = f"http://{GO_SERVICE_HOST}:{GO_SERVICE_PORT}/ping/")
    # extracting data in json format
    print(r)
    return str(f"Your response is: {r.text}")

# main driver function
if __name__ == '__main__':
	app.run(host="0.0.0.0", port=PYTHON_SERVICE_PORT)
