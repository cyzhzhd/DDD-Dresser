package com.dresser.application.products;

import com.dresser.domain.products.Product;
import com.dresser.domain.products.ProductFactory;
import com.dresser.domain.products.ProductId;
import com.dresser.domain.products.ProductRepository;
import org.springframework.stereotype.Service;

@Service
public class ProductUpdateService implements IProductUpdateApplicationService {
    
    private final ProductRepository repository;
    private final ProductFactory factory;
    
    public ProductUpdateService(ProductRepository repository, ProductFactory factory) {
        this.repository = repository;
        this.factory = factory;
    }
    
    @Override
    public ProductDTO update(ProductUpdateCommand command) {
        Product product = repository.findById(new ProductId(command.getId()));
        if (product == null || product.isDeleted()) {
            throw new RuntimeException("Product not found with id: " + command.getId());
        }
        
        product.update(
            command.getName(),
            command.getCategory(),
            command.getPrice(),
            command.getSizes()
        );
        
        Product updatedProduct = repository.save(product);
        return ProductDTO.fromDomain(updatedProduct);
    }
} 