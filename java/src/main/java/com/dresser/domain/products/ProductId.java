package com.dresser.domain.products;

public class ProductId {
    private final int value;
    
    public ProductId(int value) {
        this.value = value;
    }
    
    public int getValue() {
        return value;
    }
    
    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        ProductId productId = (ProductId) o;
        return value == productId.value;
    }
    
    @Override
    public int hashCode() {
        return Integer.hashCode(value);
    }
} 