# AI-Powered E-commerce Platform

## Overview

The AI-Powered E-commerce Platform is a full-stack application designed to provide users with a seamless online shopping experience. The platform leverages AI to deliver personalized product recommendations based on user behavior and preferences. This project includes both a backend API and a frontend user interface.

**Background**: The e-commerce industry has seen a significant transformation with the integration of AI technologies, providing personalized shopping experiences. This capstone project aims to develop an AI-powered e-commerce platform that enhances user experiences through personalized product recommendations.

**Purpose**: The purpose of this project is to create a fully functional e-commerce platform where users can browse products, add items to their cart, and make purchases. The AI integration will offer personalized product recommendations based on user interactions and preferences, thereby improving the overall shopping experience.

## Features

- **User Authentication & Authorization**
  - JWT-based user authentication
  - Secure login and registration

- **Product Management**
  - CRUD operations for products
  - View product details and list

- **Shopping Cart & Transactions**
  - Add items to the cart
  - Process purchases and manage transactions

- **AI-Powered Recommendations**
  - Personalized product recommendations based on user interactions and purchase history

## Technologies Used

- **Backend**: Go (Golang), Gin Framework, GORM, PostgreSQL
- **Frontend**: React.js, Axios
- **AI**: Integrated recommendation engine (custom logic)

## Project Structure
![Directory Structure](/DirectoryStructure.png)
Notes : This does not include the test folder or else the structure would be too complex to be displayed in one picture.

### ERD

The Entity-Relationship Diagram (ERD) for the database schema is as follows:

![ERD Diagram](/ERD-Capstone-AI.drawio.png)

## Setup Instructions

### Backend

1\. **Clone the Repository**

   ```bash

   git clone <your-repository-url>

   cd ecommerce-platform/backend

   ```

2\. **Install Dependencies**

   ```bash

   go mod tidy

   ```

3\. **Configure Environment Variables**

   Create a `.env` file in the `backend` directory with the following content:

   ```env

   DATABASE_URL=your_database_url

   JWT_SECRET=your_jwt_secret_key

   ```

4\. **Run the Application**

   ```bash

   go run main.go

   ```

### Frontend

1\. **Navigate to the Frontend Directory**

   ```bash

   cd ../frontend

   ```

2\. **Install Dependencies**

   ```bash

   npm install

   ```

3\. **Start the Development Server**

   ```bash

   npm start

   ```

**Daftar API Endpoints Backend**
--------------------------------

Berikut adalah daftar endpoint API yang tersedia pada backend aplikasi, beserta deskripsi dan format data yang diproses:

### **Autentikasi**

1.  **Register User**

    -   **URL:** `/register`
    -   **Method:** `POST`
    -   **Body:**

        json

        Copy code

        `{
          "username": "string",
          "password": "string"
        }`

    -   **Response:**

        json

        Copy code

        `{
          "message": "User registered successfully"
        }`

2.  **Login User**

    -   **URL:** `/login`
    -   **Method:** `POST`
    -   **Body:**

        json

        Copy code

        `{
          "username": "string",
          "password": "string"
        }`

    -   **Response:**

        json

        Copy code

        `{
          "token": "jwt_token"
        }`

### **Produk**

1.  **Get Products**

    -   **URL:** `/products`
    -   **Method:** `GET`
    -   **Response:**

        json

        Copy code

        `[
          {
            "id": "integer",
            "name": "string",
            "description": "string",
            "price": "decimal",
            "stock": "integer"
          }
        ]`

2.  **Create Product**

    -   **URL:** `/products`
    -   **Method:** `POST`
    -   **Body:**

        json

        Copy code

        `{
          "name": "string",
          "description": "string",
          "price": "decimal",
          "stock": "integer"
        }`

    -   **Response:**

        json

        Copy code

        `{
          "message": "Product created successfully"
        }`

### **Keranjang**

1.  **Get Cart Items**

    -   **URL:** `/cart`
    -   **Method:** `GET`
    -   **Response:**

        `[
          {
            "id": "integer",
            "name": "string",
            "quantity": "integer",
            "price": "decimal"
          }
        ]`

2.  **Add To Cart**

    -   **URL:** `/cart`
    -   **Method:** `POST`
    -   **Body:**

        `{
          "product_id": "integer",
          "quantity": "integer"
        }`

    -   **Response:**

        `{
          "message": "Item added to cart"
        }`

3.  **Remove From Cart**

    -   **URL:** `/cart/:id`
    -   **Method:** `DELETE`
    -   **Response:**

        `{
          "message": "Item removed from cart"
        }`

### **Flow Interaksi**

1.  **Registrasi dan Login:**

    -   Pengguna mengisi formulir registrasi atau login di frontend.
    -   Data dikirim ke backend melalui endpoint `/register` atau `/login`.
    -   Backend memproses data dan mengembalikan token JWT (untuk login) atau pesan sukses (untuk registrasi).
2.  **Melihat dan Mengelola Produk:**

    -   Pengguna dapat melihat daftar produk menggunakan endpoint `/products`.
    -   Pengguna dapat menambahkan produk ke keranjang menggunakan endpoint `/cart` (metode `POST`).
    -   Produk dapat dihapus dari keranjang menggunakan endpoint `/cart/:id` (metode `DELETE`).
3.  **Menambah dan Menghapus Item dari Keranjang:**

    -   Ketika pengguna menambahkan item ke keranjang, frontend mengirimkan permintaan `POST` ke endpoint `/cart`.
    -   Ketika pengguna menghapus item dari keranjang, frontend mengirimkan permintaan `DELETE` ke endpoint `/cart/:id`.
4.  **Menampilkan Keranjang:**

    -   Untuk menampilkan item dalam keranjang, frontend mengirimkan permintaan `GET` ke endpoint `/cart`.

### Backend

Run unit tests with:

```bash

go test ./... -cover

```

### Frontend

Run unit tests with:

```bash

npm test --coverage

```
