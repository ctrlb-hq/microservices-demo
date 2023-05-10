package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonMappingException;

@RestController
@RequestMapping("/numbers")
// DB PORT: 5432
public class NumberResource {
    @Autowired
    NumberService service;

    @GetMapping("/{id}")
    public ResponseEntity<NumberEntity> getNumber(@PathVariable("id") String id) throws Exception, JsonMappingException, JsonProcessingException {
        // Logic to retrieve number with given id
        NumberEntity number = service.getNumberById(id);
        if (number == null) {
            return ResponseEntity.notFound().build();
        } else if(number.getNumber() <= 0){
            throw new Exception("Cannot handle negative numbers!");
        }
        else {
            return ResponseEntity.ok(number);
        }
    }

    @PostMapping
    public ResponseEntity<NumberEntity> postNumber(@RequestBody NumberEntity number) {
        NumberEntity processedEntity = service.processNumber(number);
        return ResponseEntity.ok(service.setNumber(processedEntity));
    }
}
