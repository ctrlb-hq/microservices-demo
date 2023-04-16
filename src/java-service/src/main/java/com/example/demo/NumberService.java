package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class NumberService {
    @Autowired
    NumberRepository repo;

    public NumberEntity getNumberById(String id) {
        return repo.findById(id).orElse(null);
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
