package com.dresser.domain.brands;

import java.util.List;

public interface BrandRepository {
    Brand save(Brand brand);
    Brand findById(BrandId id);
    List<Brand> findAll();
    void delete(BrandId id);
    int nextId();
} 