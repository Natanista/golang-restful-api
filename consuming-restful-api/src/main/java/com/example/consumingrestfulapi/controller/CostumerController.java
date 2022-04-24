package com.example.consumingrestfulapi.controller;

import com.example.consumingrestfulapi.model.Costumer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import java.util.Arrays;
import java.util.List;

@RestController
@RequestMapping("/costumers")
public class CostumerController {

    @Autowired
    private RestTemplate restTemplate;

    @GetMapping
    public List<Object> getCostumers(){
        String url = "http://localhost:8080/customers";
        Object[] customers = restTemplate.getForObject(url, Object[].class);

        return Arrays.asList(customers);
    }

}
