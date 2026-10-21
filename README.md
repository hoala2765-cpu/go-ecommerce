# Go E-commerce Backend

Backend thương mại điện tử viết bằng Go.

## Tech Stack
- Go + Gin (Web Framework)
- GORM + PostgreSQL (Database)
- Docker (Containerization)

## Tính năng
- [x] CRUD Sản phẩm
- [x] Kết nối Database, Auto Migrate
- [ ] Đăng ký / Đăng nhập (JWT)
- [ ] Giỏ hàng & Đặt hàng

## Cách chạy
1. Khởi động DB: `docker run --name pg-ecommerce -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=ecommerce -p 5432:5432 -d postgres:16`
2. Cấu hình file `.env`
3. Chạy server: `go run cmd/api/main.go`