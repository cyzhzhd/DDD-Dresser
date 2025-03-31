package com.dresser.application.products;

import com.dresser.domain.products.Product;
import com.dresser.domain.products.ProductFactory;
import com.dresser.domain.products.ProductRepository;
import org.springframework.stereotype.Service;

@Service
public class ProductRegisterService implements IProductRegisterApplicationService {
    
    private final ProductRepository repository;
    private final ProductFactory factory;
    
    public ProductRegisterService(ProductRepository repository, ProductFactory factory) {
        this.repository = repository;
        this.factory = factory;
    }
    
    @Override
    public ProductDTO register(ProductRegisterCommand command) {
        Product product = factory.create(
            command.getBrandId(),
            command.getName(),
            command.getCategory(),
            command.getPrice(),
            command.getSizes()
        );
        
        Product savedProduct = repository.save(product);
        return ProductDTO.fromDomain(savedProduct);
    }
} 