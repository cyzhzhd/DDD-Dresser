package com.dresser.infrastructure.persistence.brands;

import com.dresser.domain.brands.Brand;
import com.dresser.domain.brands.BrandId;
import com.dresser.domain.brands.BrandRepository;
import com.dresser.infrastructure.db.DB;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@Repository
public class BrandRepositoryImpl implements BrandRepository {
    
    private static final String COLLECTION = "brands";
    private final DB database;
    
    public BrandRepositoryImpl(DB database) {
        this.database = database;
    }
    
    @Override
    public Brand save(Brand brand) {
        database.create(COLLECTION, brand.getId().getValue(), brand);
        return brand;
    }
    
    @Override
    public Brand findById(BrandId id) {
        return (Brand) database.findById(COLLECTION, id.getValue());
    }
    
    @Override
    public List<Brand> findAll() {
        List<Object> objects = database.findAll(COLLECTION);
        return objects.stream()
                .map(obj -> (Brand) obj)
                .collect(Collectors.toList());
    }
    
    @Override
    public void delete(BrandId id) {
        database.delete(COLLECTION, id.getValue());
    }
    
    @Override
    public int nextId() {
        return database.nextId(COLLECTION);
    }
} 