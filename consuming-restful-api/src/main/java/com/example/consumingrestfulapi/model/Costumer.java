package com.example.consumingrestfulapi.model;


import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

@Getter
@Setter
@ToString
@AllArgsConstructor
public class Costumer {

    private Long id;
    private String name;
    private String address;
    private Double balance;

}
