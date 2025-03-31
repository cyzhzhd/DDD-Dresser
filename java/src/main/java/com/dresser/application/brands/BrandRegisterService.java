package com.dresser.application.brands;

import com.dresser.domain.brands.Brand;
import com.dresser.domain.brands.BrandFactory;
import com.dresser.domain.brands.BrandRepository;
import org.springframework.stereotype.Service;

@Service
public class BrandRegisterService implements IBrandRegisterApplicationService {
    
    private final BrandRepository repository;
    private final BrandFactory factory;
    
    public BrandRegisterService(BrandRepository repository, BrandFactory factory) {
        this.repository = repository;
        this.factory = factory;
    }
    
    @Override
    public BrandDTO register(BrandRegisterCommand command) {
        Brand brand = factory.create(command.getName());
        Brand savedBrand = repository.save(brand);
        return BrandDTO.fromDomain(savedBrand);
    }
} 