package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/numbers")
// DB PORT: 5432
public class NumberResource {
    @Autowired
    NumberService service;

    @GetMapping("/{id}")
    public ResponseEntity<NumberEntity> getNumber(@PathVariable("id") String id) {
        // Logic to retrieve number with given id
        NumberEntity number = service.getNumberById(id);
        if (number == null) {
            return ResponseEntity.notFound().build();
        } else {
            return ResponseEntity.ok(number);
        }
    }

    @PostMapping
    public ResponseEntity<NumberEntity> postNumber(@RequestBody NumberEntity number) {
        NumberEntity processedEntity = service.processNumber(number);
        return ResponseEntity.ok(service.setNumber(processedEntity));
    }
}
