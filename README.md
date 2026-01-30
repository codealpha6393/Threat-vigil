# Threat-Vigil 🛡️

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Architecture](https://img.shields.io/badge/Architecture-Clean%20%2F%20Hexagonal-orange)
![License](https://img.shields.io/badge/License-MIT-blue)

A high-performance **Security Incident Logging API** built with **Go (Golang)** and **Gin**.

This project demonstrates **Clean Architecture (Ports & Adapters / Hexagonal)** principles to strictly decouple business logic from infrastructure. It is designed to simulate a scalable backend service that ingests, classifies, and stores security threat logs (e.g., SQL Injection attempts, Brute Force attacks) in real time.

---

## 🚀 Why This Architecture?

I chose **Clean Architecture (Hexagonal)** to ensure the system remains scalable and maintainable.

- **Decoupled Logic:**  
  The core business rules (`internal/core`) have zero dependencies on the HTTP framework or database.

- **Interchangeable Infrastructure:**  
  The database is accessed via an interface (`Port`). This allows swapping the current in-memory storage with **PostgreSQL** or **MongoDB** without changing business logic.

- **Testability:**  
  Separation of concerns makes it easy to write unit tests for the core logic without mocking the HTTP server.

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin (High-performance HTTP web framework)
- **Architecture:** Ports & Adapters (Clean Architecture)
- **Data Format:** JSON
- **Database:** In-Memory Store (Designed for easy migration to PostgreSQL)

---

## 📂 Project Structure

The folder structure follows the standard Go project layout:

```text
/cmd
  └── main.go              # Application Entry Point

/internal
  ├── core                 # 🧠 The "Brain" (Business Logic)
  │   ├── domain            # Data Entities (Incident Structs)
  │   ├── ports             # Interfaces (Defines HOW we talk to DB/API)
  │   └── services          # Logic (Validation, Severity Assignment)
  │
  └── adapter               # 🔌 The "Plugs" (Infrastructure)
      ├── handler           # HTTP Handlers (Gin Routes)
      └── repository        # Database Implementation (In-Memory)
⚡ API Endpoints
1️⃣ Report an Incident
POST /incidents
Ingests a new security threat log.

Request Body
{
  "type": "SQL Injection",
  "severity": "CRITICAL",
  "raw_log": "SELECT * FROM users WHERE 1=1"
}
Response (201 Created)
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "OPEN",
  "created_at": "2026-01-30T10:00:00Z"
}
2️⃣ Get All Incidents
GET /incidents
Retrieves a list of all recorded threats.

Optional Query Parameters
?severity=CRITICAL
🔧 Installation & Setup Guide
Prerequisites
Go 1.21 or higher installed

Step 1: Clone the Repository
git clone https://github.com/codealpha6393/Threat-vigil.git
cd Threat-vigil
Step 2: Install Dependencies
go mod tidy
Step 3: Run the Server
go run cmd/main.go
You should see:

[GIN-debug] Listening and serving HTTP on :8080
Step 4: Test the API
Using Postman, cURL, or PowerShell:

Invoke-RestMethod `
  -Method Post `
  -Uri "http://localhost:8080/incidents" `
  -ContentType "application/json" `
  -Body '{"type":"XSS Attack", "severity":"HIGH"}'
🔮 Roadmap
 Persistence: Migrate from In-Memory DB to PostgreSQL

 Authentication: Add JWT middleware to secure API endpoints

 Caching: Implement Redis for high-speed read operations

 Docker: Add Dockerfile and docker-compose for containerized deployment

👤 Author
Vivek

Built to explore scalable backend patterns in Go for high-security environments.


---
