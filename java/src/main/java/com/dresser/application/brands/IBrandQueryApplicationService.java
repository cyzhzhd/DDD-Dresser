package com.dresser.application.brands;

import java.util.List;

public interface IBrandQueryApplicationService {
    List<BrandDTO> getAll();
    BrandDTO get(int id);
} 