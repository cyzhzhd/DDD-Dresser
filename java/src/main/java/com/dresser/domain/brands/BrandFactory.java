package com.dresser.domain.brands;

import org.springframework.stereotype.Component;

@Component
public class BrandFactory {
    
    private final BrandRepository repository;
    
    public BrandFactory(BrandRepository repository) {
        this.repository = repository;
    }
    
    public Brand create(String name) {
        int id = repository.nextId();
        return Brand.create(id, name);
    }
} 