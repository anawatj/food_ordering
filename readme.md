# 🍱 ระบบสั่งอาหารออนไลน์ (Food Ordering System)

ระบบนี้เป็น API สำหรับการสั่งอาหารออนไลน์ พัฒนาโดยใช้ภาษา Go พร้อมสถาปัตยกรรม Clean Architecture และ Domain-Driven Design (DDD) เพื่อแยกความรับผิดชอบของแต่ละชั้นอย่างชัดเจน

---

# Description

Onedaycat Assignment By Anawat Jarusiripot

---

## 🔧 Tech Stack

- 💻 Language: Go 1.20+
- 🧩 Framework: [Gin Web Framework](https://github.com/gin-gonic/gin)
- 🏛️ Architecture: Clean Architecture + DDD
- 🧪 Testing: Go test + Testify
- 💾 In-Memory Storage (mocked storage)

---

## 📦 โครงสร้างโปรเจกต์

```plaintext
├── cmd/                  # main entry point
│   └── app/
│       └── main.go
├── internal/
│   ├── domain/           # Domain Layer (Entities, Services, Interfaces)
│   ├── infrastructure/   # Memory repositories, storage, etc.
│   └── application/      # Service interfaces, DTOs
├── tests/                # Unit tests (แยกออกจาก package จริง)
├── go.mod / go.sum
└── README.md
```
