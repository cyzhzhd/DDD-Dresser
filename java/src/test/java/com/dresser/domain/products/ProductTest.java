package com.dresser.domain.products;

import com.dresser.domain.brands.BrandId;
import org.junit.jupiter.api.Test;

import java.math.BigDecimal;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class ProductTest {
    
    @Test
    void testCreateProduct() {
        // Arrange
        int id = 1;
        int brandId = 2;
        String name = "Test Product";
        String category = "Test Category";
        BigDecimal price = new BigDecimal("99.99");
        List<String> sizes = Arrays.asList("S", "M", "L");
        
        // Act
        Product product = Product.create(id, brandId, name, category, price, sizes);
        
        // Assert
        assertNotNull(product);
        assertEquals(id, product.getId().getValue());
        assertEquals(brandId, product.getBrandId().getValue());
        assertEquals(name, product.getName());
        assertEquals(category, product.getCategory());
        assertEquals(price, product.getPrice());
        assertEquals(sizes, product.getSizes());
        assertFalse(product.isDeleted());
    }
    
    @Test
    void testUpdateProduct() {
        // Arrange
        Product product = Product.create(1, 2, "Original Name", "Original Category", 
            new BigDecimal("99.99"), Arrays.asList("S", "M"));
        
        String newName = "Updated Name";
        String newCategory = "Updated Category";
        BigDecimal newPrice = new BigDecimal("149.99");
        List<String> newSizes = Arrays.asList("S", "M", "L", "XL");
        
        // Act
        product.update(newName, newCategory, newPrice, newSizes);
        
        // Assert
        assertEquals(newName, product.getName());
        assertEquals(newCategory, product.getCategory());
        assertEquals(newPrice, product.getPrice());
        assertEquals(newSizes, product.getSizes());
    }
    
    @Test
    void testMarkAsDeleted() {
        // Arrange
        Product product = Product.create(1, 2, "Test Product", "Test Category", 
            new BigDecimal("99.99"), Arrays.asList("S", "M"));
        
        // Act
        product.markAsDeleted();
        
        // Assert
        assertTrue(product.isDeleted());
    }
} 