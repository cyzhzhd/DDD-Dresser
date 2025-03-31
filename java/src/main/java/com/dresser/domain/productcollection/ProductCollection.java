package com.dresser.domain.productcollection;

import com.dresser.domain.brands.BrandId;
import com.dresser.domain.categories.Category;
import com.dresser.domain.products.Product;

import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

public class ProductCollection {
    private final String name;
    private final List<Product> products;
    
    private ProductCollection(String name, List<Product> products) {
        if (name == null || name.trim().isEmpty()) {
            throw new IllegalArgumentException("Collection name cannot be null or empty");
        }
        
        this.name = name;
        this.products = new ArrayList<>(products);
    }
    
    public static ProductCollection of(String name, List<Product> products) {
        if (products == null) {
            products = Collections.emptyList();
        }
        return new ProductCollection(name, products);
    }
    
    public String getName() {
        return name;
    }
    
    public List<Product> getProducts() {
        return products;
    }
    
    public int getProductCount() {
        return products.size();
    }
    
    public List<String> getCategories() {
        List<String> categories = new ArrayList<>();
        for (Product product : products) {
            String category = product.getCategory();
            if (!categories.contains(category)) {
                categories.add(category);
            }
        }
        return categories;
    }
    
    public ProductCollection filterByCategory(String category) {
        if (category == null || category.trim().isEmpty()) {
            return this;
        }
        
        List<Product> filtered = products.stream()
                .filter(product -> category.equals(product.getCategory()))
                .collect(Collectors.toList());
        
        return new ProductCollection(this.name + " (Category: " + category + ")", filtered);
    }
    
    public ProductCollection filterByBrand(BrandId brandId) {
        if (brandId == null) {
            return this;
        }
        
        List<Product> filtered = products.stream()
                .filter(product -> brandId.equals(product.getBrandId()))
                .collect(Collectors.toList());
        
        return new ProductCollection(this.name + " (Brand: " + brandId.getValue() + ")", filtered);
    }
    
    public Optional<Product> getLowestPriceProduct() {
        if (products.isEmpty()) {
            return Optional.empty();
        }
        
        return products.stream()
                .filter(product -> !product.isDeleted())
                .min(Comparator.comparing(Product::getPrice));
    }
    
    public Optional<Product> getHighestPriceProduct() {
        if (products.isEmpty()) {
            return Optional.empty();
        }
        
        return products.stream()
                .filter(product -> !product.isDeleted())
                .max(Comparator.comparing(Product::getPrice));
    }
    
    public ProductCollection getLowestProductsByCategories() {
        List<Product> lowestProducts = new ArrayList<>();
        
        // Group products by category
        var productsByCategory = products.stream()
            .collect(Collectors.groupingBy(Product::getCategory));
        
        // Get lowest price product for each category
        for (var entry : productsByCategory.entrySet()) {
            entry.getValue().stream()
                .min(Comparator.comparing(Product::getPrice))
                .ifPresent(lowestProducts::add);
        }
        
        return new ProductCollection(name + " (Lowest by Category)", lowestProducts);
    }
    
    public BigDecimal getTotalPrice() {
        return products.stream()
                .filter(product -> !product.isDeleted())
                .map(Product::getPrice)
                .reduce(BigDecimal.ZERO, BigDecimal::add);
    }
} 