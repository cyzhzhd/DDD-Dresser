package com.dresser.application.products;

import com.dresser.domain.brands.BrandId;
import com.dresser.domain.products.Product;
import com.dresser.domain.products.ProductId;
import com.dresser.domain.products.ProductRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class ProductQueryService implements IProductQueryApplicationService {
    
    private final ProductRepository repository;
    
    public ProductQueryService(ProductRepository repository) {
        this.repository = repository;
    }
    
    @Override
    public List<ProductDTO> getAll() {
        return repository.findAll().stream()
                .filter(product -> !product.isDeleted())
                .map(ProductDTO::fromDomain)
                .collect(Collectors.toList());
    }
    
    @Override
    public ProductDTO get(int id) {
        Product product = repository.findById(new ProductId(id));
        if (product == null || product.isDeleted()) {
            throw new RuntimeException("Product not found with id: " + id);
        }
        return ProductDTO.fromDomain(product);
    }
    
    @Override
    public List<ProductDTO> getByBrandId(int brandId) {
        return repository.findByBrandId(new BrandId(brandId)).stream()
                .filter(product -> !product.isDeleted())
                .map(ProductDTO::fromDomain)
                .collect(Collectors.toList());
    }
    
    @Override
    public List<ProductDTO> getByCategory(String category) {
        return repository.findByCategory(category).stream()
                .filter(product -> !product.isDeleted())
                .map(ProductDTO::fromDomain)
                .collect(Collectors.toList());
    }
} 