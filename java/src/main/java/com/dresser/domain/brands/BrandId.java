package com.dresser.domain.brands;

public class BrandId {
    private final int value;
    
    public BrandId(int value) {
        this.value = value;
    }
    
    public int getValue() {
        return value;
    }
    
    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        BrandId brandId = (BrandId) o;
        return value == brandId.value;
    }
    
    @Override
    public int hashCode() {
        return Integer.hashCode(value);
    }
} 