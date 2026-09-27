# Coaching Management System API

A production-oriented backend API for managing coaching centers, branches, teachers, students, batches, subjects, examinations, and academic activities.

Built with **Go, Echo, GORM, and PostgreSQL**, following a modular domain-based architecture.

---

## Roles & Permissions

| Role             | Description                  | Key Permissions                              |
| ---------------- | ---------------------------- | -------------------------------------------- |
| **SUPER_ADMIN**  | Users who purchase medicines | Manage Full platform/coaching administration |
| **ADMIN**        | Users who purchase medicines | Manage a coaching, create bracnh             |
| **BRANCH_ADMIN** | Platform moderators          | Manage one branch                            |
| **STUDENT**      | Medicine vendors/pharmacies  | Manage Student-specific access               |
| **TEACHER**      | Medicine vendors/pharmacies  | Manage Teaching/academic operations          |
| **STAFF**        | Medicine vendors/pharmacies  | Manage Limited operational access            |

> 💡 **Note**:

## Features:

- Authentication & Authorization
- Role-Based Access Control
- Multi-Coaching Management
- Branch Management
- Teacher Management
- Student Management
- Subject Management
- Batch Management
- Examination Management
- Student Results & Submissions
- JWT-based authentication
- PostgreSQL database
- Docker support
- Domain-based modular architecture
- Request validation
- Centralized error handling

### Publice Features

### Super_Admin Features

### Admin Features

### Branch_Manger Features

### Teacher Features

### Student Features

### Staff Features

---

## Tech Stack

| Category         | Technology   |
| ---------------- | ------------ |
| Language         | Go           |
| Framework        | Echo         |
| ORM              | GORM         |
| Database         | PostgreSQL   |
| Authentication   | JWT          |
| Password Hashing | bcrypt       |
| Containerization | Docker       |
| API Testing      | Postman      |
| Version Control  | Git / GitHub |

---

## Architecture

The project follows a modular domain-based architecture.

```text
Request
   │
   ▼
Handler
   │
   ▼
Service
   │
   ▼
Repository
   │
   ▼
PostgreSQL
```

Each domain owns its business logic and data-access responsibilities.

```text
internal/
└── domain/
    ├── auth/
    ├── coaching/
    ├── branch/
    ├── teacher/
    ├── student/
    ├── batch/
    ├── subject/
    ├── exam/
    └── submission/
```

---

## Domain Structure

Each major domain follows a similar structure:

```text
coaching/
├── entity.go
├── request.go
├── response.go
├── repository.go
├── service.go
├── handler.go
├── mapper.go
└── register.go
```

### Responsibilities

**Entity**

Defines database models and relationships.

**Request / Response**

Defines the API contract.

**Repository**

Handles database operations.

**Service**

Contains business rules and application logic.

**Handler**

Handles HTTP requests and responses.

**Mapper**

Converts entities into API responses.

**Register**

Registers routes and dependencies.

---

## Authentication

The API uses JWT-based authentication.

Authentication flow:

```text
Login
  │
  ▼
Validate Credentials
  │
  ▼
Generate Access Token
  │
  ▼
Client
  │
  ▼
Authorization: Bearer <token>
  │
  ▼
JWT Middleware
  │
  ▼
Protected Route
```

Example:

```http
Authorization: Bearer <access_token>
```

JWT claims include information such as:

```json
{
  "user_id": 1,
  "coaching_id": 3,
  "role": "SUPER_ADMIN",
  "session_id": 31,
  "token_type": "access"
}
```

---

## Multi-Coaching Architecture

The system supports multiple coaching organizations.

```text
Platform
│
├── Coaching A
│   ├── Branches
│   ├── Teachers
│   ├── Students
│   └── Batches
│
└── Coaching B
    ├── Branches
    ├── Teachers
    ├── Students
    └── Batches
```

Most tenant-specific entities contain a `coaching_id` to maintain data isolation.

---

## Database Relationship

A simplified relationship:

```text
Coaching
   │
   ├── Branch
   │
   ├── Teacher
   │
   ├── Student
   │
   └── Batch
        │
        ├── Students
        └── Teachers

Subject
   │
   └── ExamSubject
            │
            └── Exam
```

---

## API Structure

Base URL:

```text
/api/v1
```

Example endpoints:

### Authentication

```http
POST /api/v1/auth/login
POST /api/v1/auth/logout
POST /api/v1/auth/refresh
```

### Coaching

```http
POST   /api/v1/coachings
GET    /api/v1/coachings
GET    /api/v1/coachings/:id
PATCH  /api/v1/coachings/:id
DELETE /api/v1/coachings/:id
```

### Branch

```http
POST /api/v1/branchs
GET  /api/v1/branchs
GET  /api/v1/branchs
PATCH  /api/v1/branches/:id
DELETE /api/v1/branchs/:id
```

### Teachers

```http
POST   /api/v1/teachers
GET    /api/v1/teachers
GET    /api/v1/teachers/:id
PATCH  /api/v1/teachers/:id
DELETE /api/v1/teachers/:id
```

### Students

```http
POST   /api/v1/students
GET    /api/v1/students
GET    /api/v1/students/:id
PATCH  /api/v1/students/:id
DELETE /api/v1/students/:id
```

---

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/your-username/coaching-management-system.git

cd coaching-management-system
```

### 2. Configure environment variables

Create a `.env` file:

```env
APP_ENV=development
APP_PORT=5000

DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=your_password
DATABASE_NAME=coaching_management

JWT_SECRET=your_secret_key
```

Never commit the real `.env` file.

Use:

```text
.env.example
```

for documenting required environment variables.

---

## 3. Install dependencies

```bash
go mod download
```

---

## 4. Run the application

```bash
go run ./cmd/server
```

The API will start on:

```text
http://localhost:5000
```

---

## Docker

Build and run:

```bash
docker compose up --build
```

Stop containers:

```bash
docker compose down
```

---

## API Testing

The API can be tested using:

- Postman
- Insomnia
- REST Client
- curl

A Postman collection can be maintained under:

```text
docs/postman/
```

---

## Security

The project follows several backend security practices:

- Password hashing with bcrypt
- JWT authentication
- Role-based authorization
- Request validation
- Protected routes
- Environment-based secrets
- Database constraints
- Centralized error handling
- Tenant-aware data access

---

## Future Improvements

- Redis caching
- Background jobs
- Email notifications
- File/image storage
- Automated tests
- API documentation with Swagger/OpenAPI
- CI/CD pipeline
- Cloud deployment
- Observability and structured logging

---

## Author

**Suvo Datta**

Backend / Full-Stack Developer

Focused on:

```text
Go
TypeScript
PostgreSQL
REST APIs
System Design
Clean Architecture
Docker
```

---

## License

This project is currently maintained as a personal/portfolio project.

## Project Overview

MediStore is a full-stack e-commerce web application for purchasing over-the-counter (OTC) medicines. Customers can browse medicines, add to cart, and place orders. Sellers manage their medicine inventory and fulfill orders. Admins oversee the platform and manage all users and listings.

---

## Roles & Permissions

| Role         | Description                  | Key Permissions                                    |
| ------------ | ---------------------------- | -------------------------------------------------- |
| **Customer** | Users who purchase medicines | Browse, cart, order, track status, leave reviews   |
| **Seller**   | Medicine vendors/pharmacies  | Manage inventory, view orders, update order status |
| **Admin**    | Platform moderators          | Manage all inventory, users, oversee orders        |

> 💡 **Note**: Users select their role during registration Admin accounts should be seeded in the database.

---

## Tech Stack

🛠️ **See [README.md](./README.md#-tech-stack) for complete technology specifications.**

---

## Features

### Public Features

- Browse all available medicines
- Search and filter by category, price, manufacturer
- View medicine details

### Customer Features

- Register and login as customer
- Add medicines to cart
- Place orders with shipping address (Cash on Delivery)
- Track order status
- Leave reviews after ordering
- Manage profile

### Seller Features

- Register and login as seller
- Add, edit, and remove medicines
- Manage stock levels
- View incoming orders
- Update order status

### Admin Features

- View all users (customers and sellers)
- Manage user status (ban/unban)
- View all medicines and orders
- Manage categories

---

## Pages & Routes

> ⚠️ **Note**: These routes are examples. You may add, edit, or remove routes based on your implementation needs.

### Public Routes

| Route       | Page             | Description                |
| ----------- | ---------------- | -------------------------- |
| `/`         | Home             | Hero, categories, featured |
| `/shop`     | Shop             | All medicines with filters |
| `/shop/:id` | Medicine Details | Info, add to cart          |
| `/login`    | Login            | Login form                 |
| `/register` | Register         | Registration form          |

### Customer Routes (Private)

| Route         | Page          | Description      |
| ------------- | ------------- | ---------------- |
| `/cart`       | Cart          | View cart items  |
| `/checkout`   | Checkout      | Shipping address |
| `/orders`     | My Orders     | Order history    |
| `/orders/:id` | Order Details | Items, status    |
| `/profile`    | Profile       | Edit info        |

### Seller Routes (Private)

| Route               | Page      | Description      |
| ------------------- | --------- | ---------------- |
| `/seller/dashboard` | Dashboard | Orders, stats    |
| `/seller/medicines` | Inventory | Manage medicines |
| `/seller/orders`    | Orders    | Update status    |

### Admin Routes (Private)

| Route               | Page       | Description       |
| ------------------- | ---------- | ----------------- |
| `/admin`            | Dashboard  | Statistics        |
| `/admin/users`      | Users      | Manage users      |
| `/admin/orders`     | Orders     | All orders        |
| `/admin/categories` | Categories | Manage categories |

---

## Database Tables

Design your own schema for the following tables:

- **Users** - Store user information and authentication details
- **Categories** - Medicine categories
- **Medicines** - Medicine/product inventory (linked to seller)
- **Orders** - Customer orders with items and status
- **Reviews** - Customer reviews for medicines

> 💡 _Think about what fields each table needs based on the features above._

---

## API Endpoints

> ⚠️ **Note**: These endpoints are examples. You may add, edit, or remove endpoints based on your implementation needs.

### Authentication

| Method | Endpoint             | Description       |
| ------ | -------------------- | ----------------- |
| POST   | `/api/auth/register` | Register new user |
| POST   | `/api/auth/login`    | Login user        |
| GET    | `/api/auth/me`       | Get current user  |

### Medicines (Public)

| Method | Endpoint             | Description                    |
| ------ | -------------------- | ------------------------------ |
| GET    | `/api/medicines`     | Get all medicines with filters |
| GET    | `/api/medicines/:id` | Get medicine details           |
| GET    | `/api/categories`    | Get all categories             |

### Orders

| Method | Endpoint          | Description       |
| ------ | ----------------- | ----------------- |
| POST   | `/api/orders`     | Create new order  |
| GET    | `/api/orders`     | Get user's orders |
| GET    | `/api/orders/:id` | Get order details |

### Seller Management

| Method | Endpoint                    | Description         |
| ------ | --------------------------- | ------------------- |
| POST   | `/api/seller/medicines`     | Add medicine        |
| PUT    | `/api/seller/medicines/:id` | Update medicine     |
| DELETE | `/api/seller/medicines/:id` | Remove medicine     |
| GET    | `/api/seller/orders`        | Get seller's orders |
| PATCH  | `/api/seller/orders/:id`    | Update order status |

### Admin

| Method | Endpoint               | Description        |
| ------ | ---------------------- | ------------------ |
| GET    | `/api/admin/users`     | Get all users      |
| PATCH  | `/api/admin/users/:id` | Update user status |

---

## Flow Diagrams

### 💊 Customer Journey

```
                              ┌──────────────┐
                              │   Register   │
                              └──────────────┘
                                     │
                                     ▼
                              ┌──────────────┐
                              │  Browse Shop │
                              └──────────────┘
                                     │
                                     ▼
                              ┌──────────────┐
                              │ Add to Cart  │
                              └──────────────┘
                                     │
                                     ▼
                              ┌──────────────┐
                              │   Checkout   │
                              └──────────────┘
                                     │
                                     ▼
                              ┌──────────────┐
                              │ Track Order  │
                              └──────────────┘
```

### 🏪 Seller Journey

```
                              ┌──────────────┐
                              │   Register   │
                              └──────────────┘
                                     │
                                     ▼
                              ┌──────────────┐
                              │Add Medicines │
                              └──────────────┘
                                     │
                                     ▼
                              ┌──────────────┐
                              │ Manage Stock │
                              └──────────────┘
                                     │
                                     ▼
                              ┌──────────────┐
                              │ View Orders  │
                              └──────────────┘
                                     │
                                     ▼
                              ┌──────────────┐
                              │Update Status │
                              └──────────────┘
```

### 📊 Order Status

```
                              ┌──────────────┐
                              │    PLACED    │
                              └──────────────┘
                               /            \
                              /              \
                        (seller)        (customer)
                        confirms         cancels
                            /                \
                           ▼                  ▼
                   ┌──────────────┐   ┌──────────────┐
                   │  PROCESSING  │   │  CANCELLED   │
                   └──────────────┘   └──────────────┘
                          │
                          ▼
                   ┌──────────────┐
                   │   SHIPPED    │
                   └──────────────┘
                          │
                          ▼
                   ┌──────────────┐
                   │  DELIVERED   │
                   └──────────────┘
```

> 💊 **Note**: OTC medicines only (no prescription required)

---

## Submission

📋 **See [README.md](./README.md) for submission guidelines, timeline, and marks.**
