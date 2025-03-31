package com.dresser.domain.products;

import com.dresser.domain.brands.BrandId;

import java.util.List;

public interface ProductRepository {
    Product save(Product product);
    Product findById(ProductId id);
    List<Product> findAll();
    List<Product> findByBrandId(BrandId brandId);
    List<Product> findByCategory(String category);
    void delete(ProductId id);
    int nextId();
} 