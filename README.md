
# 🧩 Go Gin REST API - User CRUD

REST API sederhana menggunakan Go (Golang), [Gin](https://github.com/gin-gonic/gin), dan [GORM](https://gorm.io/) untuk operasi CRUD user.

## 📦 Struktur Proyek

```
.
├── config/
│   └── database.go
├── controllers/
│   └── user_controller.go
├── models/
│   └── user.go
├── routes/
│   └── user_routes.go
├── .env
├── .env_example
├── .gitignore
├── go.mod
├── go.sum
└── main.go
```

---

## ⚙️ Setup & Menjalankan

### 1. Clone Project

```bash
git clone https://github.com/giandiport80/go-gin-api-user.git
cd go-gin-api-user
```

### 2. Buat `.env`

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=apl_go_gin_api_user
```

### 3. Jalankan

```bash
go mod tidy
go run main.go
```

---

## 📡 API Endpoint

Semua endpoint berada di bawah prefix `/api/users`

| Method | Endpoint       | Deskripsi               |
|--------|----------------|-------------------------|
| GET    | `/`            | Ambil semua user        |
| GET    | `/:id`         | Ambil user berdasarkan ID |
| POST   | `/`            | Tambah user baru        |
| PUT    | `/:id`         | Update user             |
| DELETE | `/:id`         | Hapus user              |

---

## 📥 Request & Response

### ✅ Create User

**POST** `/api/users/`

```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Data berhasil disimpan"
}
```

---

### ✅ Get All Users

**GET** `/api/users/`

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

---

### ✅ Update User

**PUT** `/api/users/1`

```json
{
  "name": "Jane Doe",
  "email": "jane@example.com"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Data berhasil disimpan"
}
```

---

### ✅ Delete User

**DELETE** `/api/users/1`

```json
{
  "success": true,
  "message": "User deleted successfully"
}
```

---

## 📚 Tools

- Gin Web Framework
- GORM ORM
- MySQL
- godotenv

---

## 🧑‍💻 Author

Made with ❤️ by **giandiport80**

---

## 📄 License

MIT License
