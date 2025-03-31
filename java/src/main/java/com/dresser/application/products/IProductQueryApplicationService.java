package com.dresser.application.products;

import java.util.List;

public interface IProductQueryApplicationService {
    List<ProductDTO> getAll();
    ProductDTO get(int id);
    List<ProductDTO> getByBrandId(int brandId);
    List<ProductDTO> getByCategory(String category);
} 