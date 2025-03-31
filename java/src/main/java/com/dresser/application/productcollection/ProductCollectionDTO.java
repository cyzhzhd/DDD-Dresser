package com.dresser.application.productcollection;

import java.util.List;

public class ProductCollectionDTO {
    private List<ProductDTO> products;
    private double totalPrice;
    
    public ProductCollectionDTO() {
    }
    
    public ProductCollectionDTO(List<ProductDTO> products, double totalPrice) {
        this.products = products;
        this.totalPrice = totalPrice;
    }
    
    // Getters
    public List<ProductDTO> getProducts() {
        return products;
    }
    
    public double getTotalPrice() {
        return totalPrice;
    }
    
    // Setters
    public void setProducts(List<ProductDTO> products) {
        this.products = products;
    }
    
    public void setTotalPrice(double totalPrice) {
        this.totalPrice = totalPrice;
    }
} 