package com.dresser.infrastructure.persistence.products;

import com.dresser.domain.brands.BrandId;
import com.dresser.domain.products.Product;
import com.dresser.domain.products.ProductId;
import com.dresser.domain.products.ProductRepository;
import com.dresser.infrastructure.db.DB;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

@Repository
public class ProductRepositoryImpl implements ProductRepository {
    
    private static final String COLLECTION = "products";
    private final DB database;
    
    public ProductRepositoryImpl(DB database) {
        this.database = database;
    }
    
    @Override
    public Product save(Product product) {
        database.create(COLLECTION, product.getId().getValue(), product);
        return product;
    }
    
    @Override
    public Product findById(ProductId id) {
        return (Product) database.findById(COLLECTION, id.getValue());
    }
    
    @Override
    public List<Product> findAll() {
        List<Object> objects = database.findAll(COLLECTION);
        return objects.stream()
                .map(obj -> (Product) obj)
                .collect(Collectors.toList());
    }
    
    @Override
    public List<Product> findByBrandId(BrandId brandId) {
        Map<String, Object> filter = new HashMap<>();
        filter.put("brandId", brandId);
        
        List<Object> objects = database.findBy(COLLECTION, filter);
        return objects.stream()
                .map(obj -> (Product) obj)
                .collect(Collectors.toList());
    }
    
    @Override
    public List<Product> findByCategory(String category) {
        Map<String, Object> filter = new HashMap<>();
        filter.put("category", category);
        
        List<Object> objects = database.findBy(COLLECTION, filter);
        return objects.stream()
                .map(obj -> (Product) obj)
                .collect(Collectors.toList());
    }
    
    @Override
    public void delete(ProductId id) {
        database.delete(COLLECTION, id.getValue());
    }
    
    @Override
    public int nextId() {
        return database.nextId(COLLECTION);
    }
} 