package com.dresser.domain.sizes;

import lombok.EqualsAndHashCode;
import lombok.Getter;

import java.util.Arrays;
import java.util.List;

@Getter
@EqualsAndHashCode
public class Size {
    private final String value;
    
    // Standard sizes
    public static final String XS = "XS";
    public static final String S = "S";
    public static final String M = "M";
    public static final String L = "L";
    public static final String XL = "XL";
    public static final String XXL = "XXL";
    
    public static final List<String> STANDARD_SIZES = 
        Arrays.asList(XS, S, M, L, XL, XXL);
    
    private Size(String value) {
        if (value == null || value.trim().isEmpty()) {
            throw new IllegalArgumentException("Size value cannot be null or empty");
        }
        this.value = value;
    }
    
    public static Size of(String value) {
        return new Size(value);
    }
    
    public static boolean isStandardSize(String value) {
        return STANDARD_SIZES.contains(value);
    }
    
    @Override
    public String toString() {
        return value;
    }
} 