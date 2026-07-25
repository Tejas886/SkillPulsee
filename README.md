# 🚀 SkillPulsee

A full-stack web application demonstrating modern DevOps practices using Docker, Docker Compose, NGINX, MySQL, GitHub Actions, and containerized backend and frontend services.

---

## 📌 Project Overview

SkillPulsee is a Dockerized full-stack application that showcases how multiple services can work together in a production-like environment.

The project includes:

- Frontend Application
- Backend API
- MySQL Database
- NGINX Reverse Proxy
- Docker Compose Orchestration
- GitHub Actions CI/CD Pipeline

---

## 🛠️ Tech Stack

### Frontend
- HTML
- CSS
- JavaScript

### Backend
- Python
- Flask

### Database
- MySQL

### DevOps
- Docker
- Docker Compose
- GitHub Actions
- NGINX

---

## 📂 Project Structure

```
SkillPulsee
│
├── .github/
│   └── workflows/
│
├── backend/
│
├── frontend/
│
├── mysql/
│
├── nginx/
│
├── docker-compose.yml
│
├── .env.example
│
└── README.md
```

---

## ⚙️ Prerequisites

Before running the project, install:

- Docker
- Docker Compose
- Git

---

## 🚀 Getting Started

### Clone Repository

```bash
git clone https://github.com/Tejas886/SkillPulsee.git
```

```bash
cd SkillPulsee
```

### Start Containers

```bash
docker compose up --build
```

### Stop Containers

```bash
docker compose down
```

---

## 🐳 Docker Services

The application consists of the following containers:

| Service | Description |
|----------|-------------|
| frontend | User Interface |
| backend | Flask REST API |
| mysql | Database |
| nginx | Reverse Proxy |

---

## 🔄 CI/CD

GitHub Actions automatically:

- Builds the application
- Validates Docker configuration
- Runs CI workflow
- Supports automated deployments

Workflow files are available inside:

```
.github/workflows/
```

---

## 🌐 Architecture

```
           User
             │
             ▼
        NGINX Reverse Proxy
          │            │
          ▼            ▼
    Frontend       Backend API
                        │
                        ▼
                    MySQL Database
```

## 👨‍💻 Author

**Tejas**

GitHub:
https://github.com/Tejas886

---

## ⭐ Features

- Dockerized Architecture
- Multi-container Application
- Reverse Proxy with NGINX
- MySQL Integration
- GitHub Actions CI/CD
- Easy Deployment using Docker Compose

---

## 📄 License

This project is intended for learning and educational purposes.
