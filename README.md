# 🚀 Valorx Interview Task - User API

This project is a simple user management API built with **Go**, using the **Gin framework** for HTTP routing and **GORM** for database interactions. It supports basic CRUD operations for users and includes authentication features such as login with email/password and **Google OAuth**.

---

## ✨ Features
- ✅ **Create, Read, Update, and Delete (CRUD)** operations for users.
- 🔒 **User authentication** with email and password.
- 🔑 **Google OAuth** for user authentication.

---

## 🏁 Getting Started

### 📌 Prerequisites
- ⚙️ **Go** 1.16 or later
- 🛢️ **PostgreSQL**
- 🔑 A **Google Cloud project** with OAuth 2.0 credentials

---

## 🛠️ Installation

### 1️⃣ Clone the repository
```bash
git clone https://github.com/syahriarreza/valorx-intv-task-01.git
cd valorx-intv-task-01
```

### 2️⃣ Set up the database
Ensure **PostgreSQL** is running and create a database for the project. Update the **.env** file with your database credentials:
```bash
DATABASE_DSN=postgres://<username>:<password>@localhost:5432/valorx_db?sslmode=disable
```

### 3️⃣ Run database migrations
Use a migration tool to apply the SQL migrations in the **migrations/** directory to set up the database schema.

### 4️⃣ Configure Google OAuth
- Create a **Google Cloud project** and set up **OAuth 2.0 credentials**.
- Update the **.env** file with your Google OAuth credentials:
```bash
OAUTH_CLIENT_ID=<your-client-id>
OAUTH_CLIENT_SECRET=<your-client-secret>
OAUTH_REDIRECT_URL=http://localhost:8080/callback
```

### 5️⃣ Run the application
```bash
go run cmd/user-service/main.go
```
The server will start on the **port specified in the `.env` file** (default is `8080`).

---

## 📡 API Endpoints

| Method  | Endpoint         | Description          |
|---------|-----------------|----------------------|
| **POST** | `/users`        | Create a new user   |
| **GET**  | `/users/:id`    | Get user by ID      |
| **PUT**  | `/users/:id`    | Update user         |
| **DELETE** | `/users/:id` | Delete user         |
| **POST** | `/login`        | Login User         |
| **GET**  | `/auth/google`  | Google OAuth Login |
| **GET**  | `/callback`     | Google OAuth Callback |

---

## 🔑 Using Google OAuth

To authenticate using **Google OAuth**, open your browser and navigate to:

👉 **[http://localhost:8080/auth/google](http://localhost:8080/auth/google)**

This will redirect you to **Google's login page**. After successful authentication, you will be redirected back to the application, where a **new user will be created** if they do not already exist.

---

## 🧪 Testing

You can use **Postman** or any other API client to test the endpoints. A **Postman collection** is provided in the `docs/postman` directory.

---

## 📜 License

This project is licensed under the **MIT License**.
