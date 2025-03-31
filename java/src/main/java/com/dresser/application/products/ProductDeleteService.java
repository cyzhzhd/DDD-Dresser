package com.dresser.application.products;

import com.dresser.domain.products.Product;
import com.dresser.domain.products.ProductId;
import com.dresser.domain.products.ProductRepository;
import org.springframework.stereotype.Service;

@Service
public class ProductDeleteService implements IProductDeleteApplicationService {
    
    private final ProductRepository repository;
    
    public ProductDeleteService(ProductRepository repository) {
        this.repository = repository;
    }
    
    @Override
    public void delete(int id) {
        Product product = repository.findById(new ProductId(id));
        if (product != null) {
            product.markAsDeleted();
            repository.save(product);
        }
    }
} 