package com.dresser.application.brands;

import com.dresser.domain.brands.BrandId;
import com.dresser.domain.brands.BrandRepository;
import org.springframework.stereotype.Service;

@Service
public class BrandDeleteService implements IBrandDeleteApplicationService {
    
    private final BrandRepository repository;
    
    public BrandDeleteService(BrandRepository repository) {
        this.repository = repository;
    }
    
    @Override
    public void delete(int id) {
        repository.delete(new BrandId(id));
    }
} 