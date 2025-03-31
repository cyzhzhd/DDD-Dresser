package com.dresser.application.brands;

import com.dresser.domain.brands.Brand;
import com.dresser.domain.brands.BrandId;
import com.dresser.domain.brands.BrandRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class BrandQueryService implements IBrandQueryApplicationService {
    
    private final BrandRepository repository;
    
    public BrandQueryService(BrandRepository repository) {
        this.repository = repository;
    }
    
    @Override
    public List<BrandDTO> getAll() {
        return repository.findAll().stream()
                .map(BrandDTO::fromDomain)
                .collect(Collectors.toList());
    }
    
    @Override
    public BrandDTO get(int id) {
        Brand brand = repository.findById(new BrandId(id));
        if (brand == null) {
            throw new RuntimeException("Brand not found with id: " + id);
        }
        return BrandDTO.fromDomain(brand);
    }
} 