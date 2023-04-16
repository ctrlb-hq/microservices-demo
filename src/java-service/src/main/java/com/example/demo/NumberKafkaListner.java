package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;

@Component
public class NumberKafkaListner {
    @Autowired
    NumberService service;
    
    @KafkaListener(topics = "${KAFKA_SERVICE_TOPIC}")
    public void onMessage(String message) throws JsonMappingException, JsonProcessingException {
        String jsonMessage = message.replace('\'', '\"');
        System.out.println("---------------------------------------------");
        System.out.println(jsonMessage);
        System.out.println("---------------------------------------------");
        ObjectMapper mapper = new ObjectMapper();
        NumberEntity number = mapper.readValue(jsonMessage, NumberEntity.class);
        NumberEntity processedEntity = service.processNumber(number);
        service.setNumber(processedEntity);
    }
}
