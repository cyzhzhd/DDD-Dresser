package com.dresser.application.productcollection;

import java.util.List;

import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.AllArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class LowestPriceProductsDTO {
    private List<ProductDTO> products;
    private double totalPrice;
} 