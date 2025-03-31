# DDD-Dresser

This is a Java implementation of the DDD-Dresser application, using Spring Boot and following Domain-Driven Design principles.

## Project Structure

The project follows a typical DDD architecture:

- **Domain Layer**: Contains the core business logic and domain models
- **Application Layer**: Contains application services that orchestrate domain operations
- **Infrastructure Layer**: Contains technical implementations like database access
- **Interfaces Layer**: Contains controllers and DTOs for the REST API

## Building and Running

### Prerequisites

- Java 17 or higher
- Maven 3.6.x or higher

### Build

```bash
mvn clean package
```

### Run

```bash
java -jar target/dresser-1.0-SNAPSHOT.jar
```

Or use Maven:

```bash
mvn spring-boot:run
```

## API Endpoints

### Brands

- `GET /api/brands` - Get all brands
- `GET /api/brands/{id}` - Get a specific brand
- `POST /api/brands` - Create a new brand
- `DELETE /api/brands/{id}` - Delete a brand

### Products

- `GET /api/products` - Get all products
- `GET /api/products?brandId={brandId}` - Get products by brand ID
- `GET /api/products?category={category}` - Get products by category
- `GET /api/products/{id}` - Get a specific product
- `POST /api/products` - Create a new product
- `PUT /api/products/{id}` - Update a product
- `DELETE /api/products/{id}` - Delete a product

### Categories

- `GET /api/categories` - Get all categories with brand information and product counts
- `GET /api/categories/lowest-products` - Get the lowest price product for each category
- `GET /api/categories/lowest-by-brand` - Get the brand with lowest total price across all categories
- `GET /api/categories/{category}/lowest-highest` - Get the lowest and highest priced products in a specific category

### Health Check

- `GET /api/health` - Get system health information and status

## Configuration

The application uses an in-memory database by default, which can be configured in `config.json` or via Spring application properties. 