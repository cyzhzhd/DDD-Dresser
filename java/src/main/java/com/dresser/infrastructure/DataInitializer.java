package com.dresser.infrastructure;

import com.dresser.application.brands.BrandRegisterCommand;
import com.dresser.application.brands.IBrandRegisterApplicationService;
import com.dresser.application.products.IProductRegisterApplicationService;
import com.dresser.application.products.ProductRegisterCommand;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.ApplicationListener;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;
import java.util.Arrays;
import java.util.List;

@Component
public class DataInitializer implements ApplicationListener<ApplicationReadyEvent> {
    
    private static final Logger logger = LoggerFactory.getLogger(DataInitializer.class);
    
    private final IBrandRegisterApplicationService brandRegisterService;
    private final IProductRegisterApplicationService productRegisterService;
    
    public DataInitializer(IBrandRegisterApplicationService brandRegisterService, 
                          IProductRegisterApplicationService productRegisterService) {
        this.brandRegisterService = brandRegisterService;
        this.productRegisterService = productRegisterService;
    }
    
    @Override
    public void onApplicationEvent(ApplicationReadyEvent event) {
        initializeData();
    }
    
    private void initializeData() {
        logger.info("Initializing sample data...");
        
        try {
            // Create sample brands
            BrandRegisterCommand nikeCommand = new BrandRegisterCommand();
            nikeCommand.setName("Nike");
            var nike = brandRegisterService.register(nikeCommand);
            
            BrandRegisterCommand adidasCommand = new BrandRegisterCommand();
            adidasCommand.setName("Adidas");
            var adidas = brandRegisterService.register(adidasCommand);
            
            BrandRegisterCommand pumaCommand = new BrandRegisterCommand();
            pumaCommand.setName("Puma");
            var puma = brandRegisterService.register(pumaCommand);
            
            // Create sample products
            createSampleProducts(nike.getId(), adidas.getId(), puma.getId());
            
            logger.info("Sample data initialization completed successfully");
        } catch (Exception e) {
            logger.error("Error initializing sample data", e);
        }
    }
    
    private void createSampleProducts(int nikeId, int adidasId, int pumaId) {
        // Nike products
        ProductRegisterCommand nikeProduct1 = new ProductRegisterCommand();
        nikeProduct1.setBrandId(nikeId);
        nikeProduct1.setName("Air Force 1");
        nikeProduct1.setCategory("Sneakers");
        nikeProduct1.setPrice(new BigDecimal("120.00"));
        nikeProduct1.setSizes(Arrays.asList("7", "8", "9", "10", "11"));
        productRegisterService.register(nikeProduct1);
        
        ProductRegisterCommand nikeProduct2 = new ProductRegisterCommand();
        nikeProduct2.setBrandId(nikeId);
        nikeProduct2.setName("Tech Fleece Hoodie");
        nikeProduct2.setCategory("Hoodies");
        nikeProduct2.setPrice(new BigDecimal("130.00"));
        nikeProduct2.setSizes(Arrays.asList("S", "M", "L", "XL"));
        productRegisterService.register(nikeProduct2);
        
        // Adidas products
        ProductRegisterCommand adidasProduct1 = new ProductRegisterCommand();
        adidasProduct1.setBrandId(adidasId);
        adidasProduct1.setName("Ultraboost 21");
        adidasProduct1.setCategory("Sneakers");
        adidasProduct1.setPrice(new BigDecimal("180.00"));
        adidasProduct1.setSizes(Arrays.asList("7", "8", "9", "10", "11", "12"));
        productRegisterService.register(adidasProduct1);
        
        ProductRegisterCommand adidasProduct2 = new ProductRegisterCommand();
        adidasProduct2.setBrandId(adidasId);
        adidasProduct2.setName("Tiro Track Pants");
        adidasProduct2.setCategory("Pants");
        adidasProduct2.setPrice(new BigDecimal("45.00"));
        adidasProduct2.setSizes(Arrays.asList("XS", "S", "M", "L", "XL"));
        productRegisterService.register(adidasProduct2);
        
        // Puma products
        ProductRegisterCommand pumaProduct1 = new ProductRegisterCommand();
        pumaProduct1.setBrandId(pumaId);
        pumaProduct1.setName("RS-X³");
        pumaProduct1.setCategory("Sneakers");
        pumaProduct1.setPrice(new BigDecimal("110.00"));
        pumaProduct1.setSizes(Arrays.asList("8", "9", "10", "11"));
        productRegisterService.register(pumaProduct1);
        
        ProductRegisterCommand pumaProduct2 = new ProductRegisterCommand();
        pumaProduct2.setBrandId(pumaId);
        pumaProduct2.setName("Essentials Logo Tee");
        pumaProduct2.setCategory("T-Shirts");
        pumaProduct2.setPrice(new BigDecimal("25.00"));
        pumaProduct2.setSizes(Arrays.asList("S", "M", "L", "XL", "XXL"));
        productRegisterService.register(pumaProduct2);
    }
} 