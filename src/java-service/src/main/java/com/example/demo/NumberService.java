package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;

@Service
public class NumberService {
    @Autowired
    NumberRepository repo;
    
    public NumberEntity getNumberById(String id) throws JsonMappingException, JsonProcessingException {
        // return repo.findById(id).orElse(null);
        String host = System.getenv("GO_SERVICE_HOST");
        String port = System.getenv("GO_SERVICE_PORT");
        String uri = "http://" + host + ":" + port + "/fetchNumber?uuid=" + id;
        RestTemplate restTemplate = new RestTemplate();
        String response = restTemplate.getForObject(uri, String.class);
        ObjectMapper mapper = new ObjectMapper();
        return mapper.readValue(response, NumberEntity.class);
    }

    public NumberEntity setNumber(NumberEntity num) {
        return repo.save(num);
    }

    public NumberEntity processNumber(NumberEntity number) {
        int num = number.getNumber();
        number.setNumber(num * num);
        return number;
    }
}
