package com.dresser.interfaces.http.handlers;

import com.dresser.application.brands.BrandDTO;
import com.dresser.application.brands.BrandRegisterCommand;
import com.dresser.application.brands.IBrandDeleteApplicationService;
import com.dresser.application.brands.IBrandQueryApplicationService;
import com.dresser.application.brands.IBrandRegisterApplicationService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/brands")
public class BrandController {
    
    private final IBrandRegisterApplicationService brandRegisterService;
    private final IBrandQueryApplicationService brandQueryService;
    private final IBrandDeleteApplicationService brandDeleteService;
    
    public BrandController(
            IBrandRegisterApplicationService brandRegisterService,
            IBrandQueryApplicationService brandQueryService,
            IBrandDeleteApplicationService brandDeleteService) {
        this.brandRegisterService = brandRegisterService;
        this.brandQueryService = brandQueryService;
        this.brandDeleteService = brandDeleteService;
    }
    
    @GetMapping
    public ResponseEntity<List<BrandDTO>> getBrands() {
        return ResponseEntity.ok(brandQueryService.getAll());
    }
    
    @GetMapping("/{id}")
    public ResponseEntity<BrandDTO> getBrand(@PathVariable int id) {
        try {
            return ResponseEntity.ok(brandQueryService.get(id));
        } catch (Exception e) {
            return ResponseEntity.notFound().build();
        }
    }
    
    @PostMapping
    public ResponseEntity<BrandDTO> createBrand(@RequestBody BrandRegisterCommand command) {
        BrandDTO brand = brandRegisterService.register(command);
        return ResponseEntity.status(HttpStatus.CREATED).body(brand);
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Map<String, Object>> deleteBrand(@PathVariable int id) {
        try {
            brandDeleteService.delete(id);
            Map<String, Object> response = new HashMap<>();
            response.put("message", "Brand deleted successfully");
            response.put("id", id);
            return ResponseEntity.ok(response);
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).build();
        }
    }
} 