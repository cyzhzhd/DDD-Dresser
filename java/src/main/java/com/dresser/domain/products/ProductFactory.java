package com.dresser.domain.products;

import com.dresser.domain.brands.BrandId;
import com.dresser.domain.brands.BrandRepository;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;
import java.util.List;

@Component
public class ProductFactory {
    
    private final ProductRepository productRepository;
    private final BrandRepository brandRepository;
    
    public ProductFactory(ProductRepository productRepository, BrandRepository brandRepository) {
        this.productRepository = productRepository;
        this.brandRepository = brandRepository;
    }
    
    public Product create(int brandId, String name, String category, BigDecimal price, List<String> sizes) {
        // Validate brand exists
        if (brandRepository.findById(new BrandId(brandId)) == null) {
            throw new IllegalArgumentException("Brand with ID " + brandId + " not found");
        }
        
        int id = productRepository.nextId();
        return Product.create(id, brandId, name, category, price, sizes);
    }
} 