package com.dresser.application.productcollection;

public class LoHiProductsByCategoryDTO {
    private String category;
    private ProductDTO lowest;
    private ProductDTO highest;
    
    public LoHiProductsByCategoryDTO() {
    }
    
    public LoHiProductsByCategoryDTO(String category, ProductDTO lowest, ProductDTO highest) {
        this.category = category;
        this.lowest = lowest;
        this.highest = highest;
    }
    
    // Getters
    public String getCategory() {
        return category;
    }
    
    public ProductDTO getLowest() {
        return lowest;
    }
    
    public ProductDTO getHighest() {
        return highest;
    }
    
    // Setters
    public void setCategory(String category) {
        this.category = category;
    }
    
    public void setLowest(ProductDTO lowest) {
        this.lowest = lowest;
    }
    
    public void setHighest(ProductDTO highest) {
        this.highest = highest;
    }
} 