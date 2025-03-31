package com.dresser.interfaces.http.handlers;

import com.dresser.application.products.IProductDeleteApplicationService;
import com.dresser.application.products.IProductQueryApplicationService;
import com.dresser.application.products.IProductRegisterApplicationService;
import com.dresser.application.products.IProductUpdateApplicationService;
import com.dresser.application.products.ProductDTO;
import com.dresser.application.products.ProductRegisterCommand;
import com.dresser.application.products.ProductUpdateCommand;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/products")
public class ProductController {
    
    private final IProductRegisterApplicationService productRegisterService;
    private final IProductQueryApplicationService productQueryService;
    private final IProductDeleteApplicationService productDeleteService;
    private final IProductUpdateApplicationService productUpdateService;
    
    public ProductController(
            IProductRegisterApplicationService productRegisterService,
            IProductQueryApplicationService productQueryService,
            IProductDeleteApplicationService productDeleteService,
            IProductUpdateApplicationService productUpdateService) {
        this.productRegisterService = productRegisterService;
        this.productQueryService = productQueryService;
        this.productDeleteService = productDeleteService;
        this.productUpdateService = productUpdateService;
    }
    
    @GetMapping
    public ResponseEntity<List<ProductDTO>> getProducts(@RequestParam(required = false) Integer brandId,
                                                      @RequestParam(required = false) String category) {
        if (brandId != null) {
            return ResponseEntity.ok(productQueryService.getByBrandId(brandId));
        } else if (category != null && !category.isEmpty()) {
            return ResponseEntity.ok(productQueryService.getByCategory(category));
        } else {
            return ResponseEntity.ok(productQueryService.getAll());
        }
    }
    
    @GetMapping("/{id}")
    public ResponseEntity<ProductDTO> getProduct(@PathVariable int id) {
        try {
            return ResponseEntity.ok(productQueryService.get(id));
        } catch (Exception e) {
            return ResponseEntity.notFound().build();
        }
    }
    
    @PostMapping
    public ResponseEntity<ProductDTO> createProduct(@RequestBody ProductRegisterCommand command) {
        try {
            ProductDTO product = productRegisterService.register(command);
            return ResponseEntity.status(HttpStatus.CREATED).body(product);
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).build();
        }
    }
    
    @PutMapping("/{id}")
    public ResponseEntity<ProductDTO> updateProduct(@PathVariable int id, @RequestBody ProductUpdateCommand command) {
        try {
            command.setId(id); // Ensure the ID in the path matches the command
            ProductDTO product = productUpdateService.update(command);
            return ResponseEntity.ok(product);
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).build();
        }
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Map<String, Object>> deleteProduct(@PathVariable int id) {
        try {
            productDeleteService.delete(id);
            Map<String, Object> response = new HashMap<>();
            response.put("message", "Product deleted successfully");
            response.put("id", id);
            return ResponseEntity.ok(response);
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).build();
        }
    }
} 