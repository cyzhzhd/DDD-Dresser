package com.dresser.application.productcollection;

import java.util.List;

public class LowestProductsByBrandDTO {
    private LowestBrandGroupDTO lowest;
    
    public LowestProductsByBrandDTO() {
    }
    
    public LowestProductsByBrandDTO(LowestBrandGroupDTO lowest) {
        this.lowest = lowest;
    }
    
    public LowestBrandGroupDTO getLowest() {
        return lowest;
    }
    
    public void setLowest(LowestBrandGroupDTO lowest) {
        this.lowest = lowest;
    }

    public static class LowestBrandGroupDTO {
        private String brand;
        private List<String> categories;
        private double totalPrice;
        
        public LowestBrandGroupDTO() {
        }
        
        public LowestBrandGroupDTO(String brand, List<String> categories, double totalPrice) {
            this.brand = brand;
            this.categories = categories;
            this.totalPrice = totalPrice;
        }
        
        public String getBrand() {
            return brand;
        }
        
        public void setBrand(String brand) {
            this.brand = brand;
        }
        
        public List<String> getCategories() {
            return categories;
        }
        
        public void setCategories(List<String> categories) {
            this.categories = categories;
        }
        
        public double getTotalPrice() {
            return totalPrice;
        }
        
        public void setTotalPrice(double totalPrice) {
            this.totalPrice = totalPrice;
        }
    }
} 