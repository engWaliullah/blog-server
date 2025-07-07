# 📝 Go Blog Server

A simple and scalable blog server built with **Golang**, **PostgreSQL**, and **Gorilla Mux** following a clean folder structure.

---

## 🚀 Features

- RESTful API for blog posts
- PostgreSQL database connection
- Environment variable support with `.env`
- Organized folder structure (handlers, models, routes, utils)
- JSON response helpers

---

## 🧱 Project Structure


---

## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/engWaliullah/blog-server
cd blog-server


### 2. Install Dependencies
```bash
go mod tidy


### DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=blogdb
DB_HOST=localhost
DB_PORT=5432



### 4. Create PostgreSQL Table
 tidyCREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL
);


### 5. Run the Server
```bash
go run cmd/main.go




