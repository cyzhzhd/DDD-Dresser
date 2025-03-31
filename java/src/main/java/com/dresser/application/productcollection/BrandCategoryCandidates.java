package com.dresser.application.productcollection;

import java.util.List;

public class BrandCategoryCandidates {
    private final String brandName;
    private final List<String> products;
    private final double totalPrice;
    
    public BrandCategoryCandidates(String brandName, List<String> products, double totalPrice) {
        this.brandName = brandName;
        this.products = products;
        this.totalPrice = totalPrice;
    }
    
    public String getBrandName() {
        return brandName;
    }
    
    public List<String> getProducts() {
        return products;
    }
    
    public double getTotalPrice() {
        return totalPrice;
    }
} 