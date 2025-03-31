package com.dresser.domain.brands;

public class Brand {
    private final BrandId id;
    private String name;
    
    private Brand(BrandId id, String name) {
        this.id = id;
        this.name = name;
    }
    
    public static Brand create(int id, String name) {
        return new Brand(new BrandId(id), name);
    }
    
    public void update(String name) {
        this.name = name;
    }
    
    // Getters
    public BrandId getId() {
        return id;
    }
    
    public String getName() {
        return name;
    }
    
    // Setters
    public void setName(String name) {
        this.name = name;
    }
    
    // For equality checking
    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Brand brand = (Brand) o;
        return id.equals(brand.id);
    }
    
    @Override
    public int hashCode() {
        return id.hashCode();
    }
} 