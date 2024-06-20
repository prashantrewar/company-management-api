# Company Management API


This is a basic API built with the Go/Golang programming language to practice backend development with Go/Golang and  PostgreSQL databases (in this case PostgreSQL).


**Technologies used**

- Go (Golang): Used as the primary programming language for backend development.
- Gin Framework: A lightweight web framework for building APIs in Go. Simplifies routing, middleware management, and request handling.

- PostgreSQL: An open-source relational database used to store and manage application data.
- GORM (Go Object Relational Mapper): A Go ORM library for interacting with databases.
- JWT (JSON Web Token): Used for secure authentication and authorization.
- HTML, CSS, JavaScript: HTML for structuring web pages, CSS for styling, and JavaScript for client-side scripting and interacting with the backend API.

- Fetch API: Handles Create and Read operations and retrieves data to update the frontend interface dynamically.


#### Directory structure

```
company-management- api/
├── main.go
├── handlers/
│   ├── auth_handler.go
│   ├── user_handler.go
│   ├── customer_handler.go
│   ├── billing_handler.go
│   ├── payroll_handler.go
├── middleware/
│   ├── auth_middleware.go
├── models/
│   ├── user.go
│   ├── customer.go
│   ├── billing.go
│   ├── payroll.go
├── db/
│   ├── postgres.go
├── utils/
│   ├── jwt.go
├── config/
│   ├── config.go
frontend/
├── index.html
├── style.css
├── script.js
.env
README.md


```

## About 

API Backend and Frontend for Company Management System

# Backend Overview

The backend for this company management system is built using Go and the Gin framework, with PostgreSQL as the database. It is structured to handle various roles such as Admin, HR, Sales, and Accountant, each with specific permissions and functionalities. The key components of the backend include:

- Authentication: Users authenticate using a JWT-based mechanism. The auth_handler.go handles login requests, verifying credentials, and generating JWT tokens.

- Role-based Access Control: Middleware checks user roles to grant access to specific endpoints. For example, only Admins can create new users, while HR can manage payroll records.

- Database Interactions: The gorm ORM is used to interact with the PostgreSQL database. Models such as User and Payroll are defined to represent database tables.

- API Endpoints:

    - /login: Authenticates users and provides a JWT token.
    - /users: Manages user creation and retrieval, restricted to Admins.
    - /payroll: Handles payroll creation and retrieval, restricted to HR and Accountants.
    - /customers and /billings: Managed by Sales role, for customer and billing operations.


# Frontend Overview

The frontend is a simple HTML and JavaScript application that interacts with the backend API. It provides a user-friendly interface for different roles to perform their respective tasks. Key features include:

- Login Page: Users login using their credentials. Upon successful login, they receive a JWT token stored in localStorage

- Role-specific Menus: After login, users are redirected to a menu page tailored to their role, displaying options relevant to their permissions.

- Dynamic Forms and Actions: Forms are provided for creating users, payrolls, customers, and billings. JavaScript handles form submissions and fetches data using the Fetch API.

- AJAX Requests: Asynchronous requests are made to the backend API to perform CRUD operations. Responses are handled to display success or error messages to the user.



## Compiling and running the server

### For Backend

the curl commands to test the API manually. These commands assume that the backend server is running on http://localhost:8080.

- Login to get JWT token

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "username": "admin",
    "password": "password"
}'

```

- Create a new user (Admin only)

```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <ADMIN_TOKEN>" \
-d '{
    "username": "john",
    "password": "password",
    "role": "HR"
}'

```

- Get all users (Admin only)

```bash
curl -X GET http://localhost:8080/users \
-H "Authorization: Bearer <ADMIN_TOKEN>"

```

- Create a payroll (HR only)

```bash
curl -X POST http://localhost:8080/payroll \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <HR_TOKEN>" \
-d '{
    "employee_name": "Jane Doe",
    "amount": 5000,
    "status": "Paid"
}'

```

- Get all payrolls (HR and Accountant only)

```bash
curl -X GET http://localhost:8080/payroll \
-H "Authorization: Bearer <HR_OR_ACCOUNTANT_TOKEN>"

```

- Create a customer (Sales only)

```bash
curl -X POST http://localhost:8080/customers \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <SALES_TOKEN>" \
-d '{
    "name": "ABC Corp",
    "address": "123 Main St",
    "email": "contact@abccorp.com"
}'

```

- Get all customers (Sales only)

```bash
curl -X GET http://localhost:8080/customers \
-H "Authorization: Bearer <SALES_TOKEN>"

```

- Create a billing (Sales only)

```bash
curl -X POST http://localhost:8080/billings \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <SALES_TOKEN>" \
-d '{
    "customer_name": Rama,
    "amount": 1000,
    "status": "Pending"
}'

```

- Get all billings (Sales and Accountant only)

```bash
curl -X GET http://localhost:8080/billings \
-H "Authorization: Bearer <SALES_OR_ACCOUNTANT_TOKEN>"

```


### For Frontend

The frontend server accessible at (http://localhost:8080)
