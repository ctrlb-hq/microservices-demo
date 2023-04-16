FROM python:3.8

ENV PYTHON_SERVICE_HOST=0.0.0.0
ENV PYTHON_SERVICE_PORT="30000"
ENV JAVA_SERVICE_HOST=0.0.0.0
ENV JAVA_SERVICE_PORT="30001"
ENV GO_SERVICE_HOST=0.0.0.0
ENV GO_SERVICE_PORT="30002"
ENV DB_SERVICE_HOST=0.0.0.0
ENV DB_SERVICE_PORT="5432"
ENV KAFKA_SERVICE_HOST=0.0.0.0
ENV KAFKA_SERVICE_PORT="9092"
ENV KAFKA_SERVICE_TOPIC=test
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=mysecretpassword
ENV DB_SERVICE_NAME=numbers

# Install Unix packages
RUN apt-get update && \
    apt-get -qq -y install curl && \
    apt-get -y install libglib2.0-0 libsm6 libxext6 libxrender-dev

# Python requirements
RUN mkdir /app
ADD requirements.txt /app/requirements.txt
RUN pip install -r /app/requirements.txt

# Copy code
ADD . /app/
RUN chmod -R 0644 /app
WORKDIR /app/
ENV PYTHONPATH=${PYTHONPATH}:.

EXPOSE 30000
EXPOSE 9092
EXPOSE 30001

ENTRYPOINT ["python", "app.py"]