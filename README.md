# 🚀 TaskFlow – Full Stack Task Management System

## 📌 Overview

TaskFlow is a full-stack task management application that allows users to:

- Register and log in securely
- Create and manage projects
- Add and manage tasks
- Perform authenticated operations using JWT

The system is fully containerized using Docker, enabling seamless local setup and consistent execution.

---

## 🛠 Tech Stack

### Backend
- Go (net/http)
- PostgreSQL
- JWT Authentication
- bcrypt (password hashing)

### Frontend
- React (Create React App)
- Axios (API communication)

### DevOps
- Docker
- Docker Compose

---

## 🧠 Architecture Overview

The system follows a **layered architecture**:


Client (React UI)
↓
API Layer (Go Handlers)
↓
Business Logic (Handlers + Middleware)
↓
Database Layer (PostgreSQL)


---

## 🔄 Application Flow

### 1. User Authentication

- User logs in via frontend
- Request sent to `/auth/login`
- Backend:
  - Verifies credentials using bcrypt
  - Generates JWT token
- Token stored in browser (localStorage)

---

### 2. Authenticated Requests

- Frontend attaches token via Axios interceptor:


Authorization: Bearer <token>


- Backend middleware:
  - Validates token
  - Allows access to protected routes

---

### 3. Data Flow


Frontend → API → Middleware → Handler → DB → Response → UI


---

## 📂 Project Structure


taskflow/
├── backend/
│ ├── cmd/main.go
│ ├── internal/
│ │ ├── handlers/
│ │ ├── middleware/
│ │ ├── db/
│ │ ├── models/
│ ├── migrations/
│
├── frontend/
│ ├── src/
│ │ ├── pages/
│ │ ├── services/
│ │ ├── App.js
│ ├── public/index.html
│
├── docker-compose.yml
├── README.md


---

## 🐳 Running the Application

### 1. Clone the repository

```bash
git clone https://github.com/diyaitis/taskflow---Diya
cd taskflow---Diya
2. Set environment variables
cp .env.example .env
3. Run the application
docker compose up --build
4. Access the app
Frontend → http://localhost:3000
Backend → http://localhost:8080
🔐 Authentication Details
JWT-based authentication
Token stored in localStorage
Protected routes secured via middleware
🧪 Test Credentials
Email: test@example.com
Password: password123
⚖️ Design Decisions
1. Go (Backend)
Lightweight and fast
Suitable for building scalable APIs
Minimal framework used for clarity
2. JWT Authentication
Stateless → scalable
Easy to integrate with frontend
3. PostgreSQL
Strong relational support
ACID compliance
Suitable for structured data
4. Dockerized Architecture
One command setup
Consistent environment across systems
5. Axios Interceptors
Centralized API handling
Automatic token injection
Global error handling
⚠️ Tradeoffs
Minimal UI (focus on backend correctness)
No pagination implemented
Limited validation on frontend
🚀 Future Improvements
Add pagination and filtering
Implement role-based access control
Add real-time updates (WebSockets)
Improve UI/UX with component libraries
Add automated testing (unit + integration)
🧠 Challenges Faced
Docker build issues and dependency management
Debugging frontend-backend communication
Handling authentication flow correctly
Ensuring proper container networking
🎯 Conclusion

TaskFlow demonstrates a clean full-stack implementation with:

Secure authentication
Modular backend design
Containerized deployment
Clear separation of concerns