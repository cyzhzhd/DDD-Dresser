package com.dresser.interfaces.http.handlers;

import com.dresser.application.productcollection.CategoryDTO;
import com.dresser.application.productcollection.ICategoriesQueryApplicationService;
import com.dresser.application.productcollection.LoHiProductsByCategoryDTO;
import com.dresser.application.productcollection.LowestProductsByBrandDTO;
import com.dresser.application.productcollection.ProductCollectionDTO;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/categories")
public class CategoriesController {
    
    private final ICategoriesQueryApplicationService categoriesQueryService;
    
    public CategoriesController(ICategoriesQueryApplicationService categoriesQueryService) {
        this.categoriesQueryService = categoriesQueryService;
    }
    
    @GetMapping
    public ResponseEntity<List<CategoryDTO>> getCategories() {
        return ResponseEntity.ok(categoriesQueryService.getAll());
    }
    
    @GetMapping("/lowest-products")
    public ResponseEntity<ProductCollectionDTO> getLowestProducts() {
        return ResponseEntity.ok(categoriesQueryService.getLowestProducts());
    }
    
    @GetMapping("/lowest-by-brand")
    public ResponseEntity<LowestProductsByBrandDTO> getLowestProductsByBrand() {
        return ResponseEntity.ok(categoriesQueryService.getLowestProductsByBrand());
    }
    
    @GetMapping("/{category}/lowest-highest")
    public ResponseEntity<LoHiProductsByCategoryDTO> getLowestAndHighestByCategory(@PathVariable String category) {
        return ResponseEntity.ok(categoriesQueryService.getLowestAndHighestProductsByCategory(category));
    }
} 