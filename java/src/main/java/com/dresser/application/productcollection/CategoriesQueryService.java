package com.dresser.application.productcollection;

import com.dresser.domain.brands.Brand;
import com.dresser.domain.brands.BrandId;
import com.dresser.domain.brands.BrandRepository;
import com.dresser.domain.categories.Category;
import com.dresser.domain.productcollection.ProductCollection;
import com.dresser.domain.products.Product;
import com.dresser.domain.products.ProductRepository;
import com.dresser.interfaces.http.GlobalExceptionHandler.ResourceNotFoundException;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.*;
import java.util.stream.Collectors;

@Service
public class CategoriesQueryService implements ICategoriesQueryApplicationService {
    
    private final BrandRepository brandRepository;
    private final ProductRepository productRepository;
    
    public CategoriesQueryService(BrandRepository brandRepository, ProductRepository productRepository) {
        this.brandRepository = brandRepository;
        this.productRepository = productRepository;
    }
    
    @Override
    public List<CategoryDTO> getAll() {
        // Get all products that are not deleted
        List<Product> products = productRepository.findAll().stream()
                .filter(product -> !product.isDeleted())
                .collect(Collectors.toList());
        
        // Create a map of category names to the products in that category
        Map<String, List<Product>> categoriesMap = new HashMap<>();
        for (Product product : products) {
            categoriesMap.computeIfAbsent(product.getCategory(), k -> new ArrayList<>()).add(product);
        }
        
        // Get all brands for lookup
        Map<Integer, Brand> brandsMap = brandRepository.findAll().stream()
                .collect(Collectors.toMap(brand -> brand.getId().getValue(), brand -> brand));
        
        // Create DTOs
        List<CategoryDTO> categories = new ArrayList<>();
        for (Map.Entry<String, List<Product>> entry : categoriesMap.entrySet()) {
            String categoryName = entry.getKey();
            List<Product> categoryProducts = entry.getValue();
            
            // Get unique brand names for this category
            Set<String> brandNames = new HashSet<>();
            for (Product product : categoryProducts) {
                Brand brand = brandsMap.get(product.getBrandId().getValue());
                if (brand != null) {
                    brandNames.add(brand.getName());
                }
            }
            
            CategoryDTO categoryDTO = new CategoryDTO();
            categoryDTO.setName(categoryName);
            categoryDTO.setBrands(new ArrayList<>(brandNames));
            categoryDTO.setCount(categoryProducts.size());
            categories.add(categoryDTO);
        }
        
        return categories;
    }

    @Override
    public ProductCollectionDTO getLowestProducts() {
        // 모든 제품 목록 가져오기
        List<Product> products = productRepository.findAll().stream()
                .filter(product -> !product.isDeleted())
                .collect(Collectors.toList());
        
        // 제품 컬렉션 생성
        ProductCollection productCollection = ProductCollection.of("All Products", products);
        
        // 카테고리별 최저가 제품 가져오기
        ProductCollection lowestProducts = productCollection.getLowestProductsByCategories();
        
        // DTO로 변환
        List<ProductDTO> productDTOs = new ArrayList<>();
        double totalPrice = 0.0;
        
        for (Product product : lowestProducts.getProducts()) {
            Brand brand = brandRepository.findById(product.getBrandId());
            if (brand == null) {
                continue;
            }
            
            ProductDTO productDTO = new ProductDTO();
            productDTO.setId(String.valueOf(product.getId().getValue()));
            productDTO.setName(product.getName());
            productDTO.setBrand(brand.getName());
            productDTO.setCategory(product.getCategory());
            productDTO.setPrice(product.getPrice().doubleValue());
            productDTO.setSizes(product.getSizes());
            
            productDTOs.add(productDTO);
            
            totalPrice += product.getPrice().doubleValue();
        }
        
        ProductCollectionDTO result = new ProductCollectionDTO();
        result.setProducts(productDTOs);
        result.setTotalPrice(totalPrice);
        return result;
    }

    @Override
    public LowestProductsByBrandDTO getLowestProductsByBrand() {
        // 모든 제품 목록 가져오기
        List<Product> products = productRepository.findAll().stream()
                .filter(product -> !product.isDeleted())
                .collect(Collectors.toList());
        
        // 제품 컬렉션 생성
        ProductCollection productCollection = ProductCollection.of("All Products", products);
        
        // 모든 브랜드 목록 가져오기
        List<Brand> brands = brandRepository.findAll();
        
        // 각 브랜드별 카테고리 최저가 제품 계산
        List<BrandCategoryCandidates> brandCandidates = new ArrayList<>();
        
        for (Brand brand : brands) {
            // 해당 브랜드의 제품만 필터링
            ProductCollection brandProducts = productCollection.filterByBrand(brand.getId());
            
            // 카테고리별 최저가 제품 가져오기
            ProductCollection lowestCategoryProducts = brandProducts.getLowestProductsByCategories();
            
            // 제품이 없는 경우 스킵
            if (lowestCategoryProducts.getProducts().isEmpty()) {
                continue;
            }
            
            // 카테고리 목록 생성
            List<String> categories = lowestCategoryProducts.getProducts().stream()
                .map(Product::getCategory)
                .collect(Collectors.toList());
            
            // 총 가격 계산
            double totalPrice = lowestCategoryProducts.getTotalPrice().doubleValue();
            
            brandCandidates.add(new BrandCategoryCandidates(
                brand.getName(),
                categories,
                totalPrice
            ));
        }
        
        // 총 가격이 가장 낮은 브랜드 찾기
        if (brandCandidates.isEmpty()) {
            throw new ResourceNotFoundException("No brands with complete product categories found");
        }
        
        BrandCategoryCandidates lowestBrand = brandCandidates.stream()
                .min(Comparator.comparing(candidate -> candidate.getTotalPrice()))
                .orElseThrow(() -> new RuntimeException("Could not determine lowest price brand"));
        
        LowestProductsByBrandDTO result = new LowestProductsByBrandDTO();
        LowestProductsByBrandDTO.LowestBrandGroupDTO groupDTO = new LowestProductsByBrandDTO.LowestBrandGroupDTO();
        groupDTO.setBrand(lowestBrand.getBrandName());
        groupDTO.setCategories(lowestBrand.getProducts());
        groupDTO.setTotalPrice(lowestBrand.getTotalPrice());
        result.setLowest(groupDTO);
        return result;
    }

    @Override
    public LoHiProductsByCategoryDTO getLowestAndHighestProductsByCategory(String category) {
        // 카테고리 유효성 검사
        if (!Category.isValidCategory(category)) {
            throw new IllegalArgumentException("Invalid category: " + category);
        }
        
        // 모든 제품 목록 가져오기
        List<Product> products = productRepository.findAll().stream()
                .filter(product -> !product.isDeleted())
                .collect(Collectors.toList());
        
        // 해당 카테고리의 제품만 필터링
        ProductCollection categoryProducts = ProductCollection.of("Category: " + category, products)
                .filterByCategory(category);
        
        // 최저가 및 최고가 제품 찾기
        Optional<Product> lowestPriceProduct = categoryProducts.getLowestPriceProduct();
        Optional<Product> highestPriceProduct = categoryProducts.getHighestPriceProduct();
        
        if (lowestPriceProduct.isEmpty() || highestPriceProduct.isEmpty()) {
            throw new ResourceNotFoundException("No products found for category: " + category);
        }
        
        // 브랜드 정보 가져오기
        Brand lowestBrand = brandRepository.findById(lowestPriceProduct.get().getBrandId());
        Brand highestBrand = brandRepository.findById(highestPriceProduct.get().getBrandId());
        
        if (lowestBrand == null || highestBrand == null) {
            throw new ResourceNotFoundException("Brand not found for product");
        }
        
        // DTO로 변환
        ProductDTO lowestDTO = new ProductDTO();
        lowestDTO.setId(String.valueOf(lowestPriceProduct.get().getId().getValue()));
        lowestDTO.setName(lowestPriceProduct.get().getName());
        lowestDTO.setBrand(lowestBrand.getName());
        lowestDTO.setCategory(category);
        lowestDTO.setPrice(lowestPriceProduct.get().getPrice().doubleValue());
        lowestDTO.setSizes(lowestPriceProduct.get().getSizes());
        
        ProductDTO highestDTO = new ProductDTO();
        highestDTO.setId(String.valueOf(highestPriceProduct.get().getId().getValue()));
        highestDTO.setName(highestPriceProduct.get().getName());
        highestDTO.setBrand(highestBrand.getName());
        highestDTO.setCategory(category);
        highestDTO.setPrice(highestPriceProduct.get().getPrice().doubleValue());
        highestDTO.setSizes(highestPriceProduct.get().getSizes());
        
        LoHiProductsByCategoryDTO result = new LoHiProductsByCategoryDTO();
        result.setCategory(category);
        result.setLowest(lowestDTO);
        result.setHighest(highestDTO);
        return result;
    }
} 