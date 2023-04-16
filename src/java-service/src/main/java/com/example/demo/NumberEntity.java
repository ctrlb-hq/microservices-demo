package com.example.demo;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "numbers")
// TODO - Adarsh get this table name in a better way from env variables
public class NumberEntity {
    @Id
    private String uuid;
    private Integer number;
    
    public String getUuid() {
        return uuid;
    }
    
    public void setUuid(String id) {
        this.uuid = id;
    }
    
    public Integer getNumber() {
        return number;
    }
    
    public void setNumber(Integer number) {
        this.number = number;
    }
}
