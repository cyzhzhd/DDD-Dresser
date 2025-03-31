package com.dresser.application.productcollection;

import java.util.List;

public interface ICategoriesQueryApplicationService {
    List<CategoryDTO> getAll();
    
    // 추가 메서드
    ProductCollectionDTO getLowestProducts();
    LowestProductsByBrandDTO getLowestProductsByBrand();
    LoHiProductsByCategoryDTO getLowestAndHighestProductsByCategory(String category);
} 