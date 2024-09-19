# AI-Powered E-commerce Platform

## Overview

The AI-Powered E-commerce Platform is a full-stack application designed to provide users with a seamless online shopping experience. The platform leverages AI to deliver personalized product recommendations based on user behavior and preferences. This project includes both a backend API and a frontend user interface.

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


## Database Setup
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

## API Endpoints

- **User Authentication**

  - `POST /register`: Register a new user

  - `POST /login`: Authenticate and get JWT token

- **Products**

  - `GET /products`: List all products

  - `POST /products`: Create a new product

- **Cart**

  - `GET /cart`: Get cart items

  - `POST /cart`: Add item to cart

  - `DELETE /cart/:id`: Remove item from cart

## Testing

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
