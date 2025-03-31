package com.dresser.domain.categories;

import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

public enum Category {
    SNEAKERS("Sneakers"),
    SHIRTS("T-Shirts"),
    PANTS("Pants"),
    HOODIES("Hoodies"),
    JACKETS("Jackets");
    
    private final String value;
    
    private static final Set<String> VALID_CATEGORIES = new HashSet<>();
    
    static {
        for (Category c : Category.values()) {
            VALID_CATEGORIES.add(c.value);
        }
    }
    
    Category(String value) {
        this.value = value;
    }
    
    public String getValue() {
        return value;
    }
    
    public static boolean isValidCategory(String category) {
        return VALID_CATEGORIES.contains(category);
    }
    
    public static Category fromString(String category) {
        if (!isValidCategory(category)) {
            throw new IllegalArgumentException("Invalid category: " + category);
        }
        return Category.valueOf(category.toUpperCase());
    }
} 