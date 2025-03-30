# DDD-Dresser

DDD-Dresser는 Domain-Driven Design principles을 적용해 만든 코디 완성 서비스입니다.  

## 구현 범위

- **도메인 모델**: 브랜드(Brands), 제품(Products), 가격(Prices), 사이즈(Sizes), 카테고리(Categories) 관리
- **계층형 아키텍처**: DDD 계층형 구조 (도메인, 응용, 인프라, 인터페이스)
- **API 기능**:
  - 브랜드: 등록, 조회, 삭제
  - 제품: 등록, 조회, 수정, 삭제
  - 코디 추천
    - 카테고리 별 최저가격 브랜드와 상품 가격, 총액을 조회
    - 단일 브랜드로 모든 카테고리 상품을 구매할 때 최저가격에 판매하는 브랜드와 카테고리의 상품가격, 총액을 조회
    - 카테고리 이름으로 최저, 최고 가격 브랜드와 상품 가격을 조회

## 코드 빌드, 테스트, 실행 방법

### 필수 조건

- Go 1.18 이상

### 실행 방법

```bash
make run
```

기본적으로 메모리 DB를 사용하며 8080 포트에서 HTTP 서버가 시작됩니다.

### 자동화 테스트 방법

```bash
make test
```

### curl 테스트 방법
mock data 로드 

```bash
make load_mock
```

카테고리 별 최저가격 브랜드와 상품 가격, 총액을 조회

```bash
curl -X GET http://localhost:8080/api/categories/lowest
```


단일 브랜드로 모든 카테고리 상품을 구매할 때 최저가격에 판매하는 브랜드와 카테고리의 상품가격, 총액을 조회

```bash
curl -X GET http://localhost:8080/api/categories/lowest/brand
```


카테고리 이름으로 최저, 최고 가격 브랜드와 상품 가격을 조회

```bash
curl -X GET http://localhost:8080/api/categories/lohi/SOCKS
```


브랜드 생성

```bash
curl -X POST http://localhost:8080/api/brands \
-d '{
  "name": "A"
}'
```

브랜드 조회

```bash
curl -X GET http://localhost:8080/api/brands \
curl -X GET http://localhost:8080/api/brands/1
```

브랜드 삭제

```bash
curl -X DELETE http://localhost:8080/api/brands/1
```


프로덕트 생성

```bash
curl -X POST http://localhost:8080/api/products \
-d '{
  "name": "adizero",
  "brand_id": 1,
  "category": "SNEAKERS",
  "amount": 100, 
  "currency": "KRW",
  "size": "220"
}'
```


프로덕트 삭제 

```bash
curl -X DELETE http://localhost:8080/api/products/1
```

## 기타 추가 정보

- **설정**: `config.json` 파일에서 데이터베이스 설정 가능
- **아키텍처**:
  - `internal/domain`: 핵심 비즈니스 로직 및 엔티티
  - `internal/application`: 유스케이스 구현
  - `internal/infrastructure`: 데이터베이스, 외부 서비스 연동
  - `internal/interface`: HTTP API 엔드포인트
