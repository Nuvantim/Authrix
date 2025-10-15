# Authrix 🔐
Authrix is a Golang-based JWT Authorization REST API designed to efficiently manage user authentication and authorization processes with high security and scalability. It supports robust features such as multi-role assignments, fine-grained permissions, secure email verification, and OTP code validation to ensure reliable access control in modern applications.

## Feature 🚀
- [x] Multi-Role and Permission each user
- [x] JWT support using RS256 algorithm
- [x] Secure validation and evolving code for protection
- [x] Provides secure email verification and one-time password (OTP) authentication.
  
## Concept 💡
<a href="https://ibb.co.com/TBKc8qTZ"><img src="https://i.ibb.co.com/MD8nV5hb/Concept-Diagram.png" alt="Concept-Diagram" border="0"></a>

## Overview 🔮
### 1. Assign Roles & Access Permissions to Each Endpoint
```go
/*
Middleware for JWT authentication across all endpoints
*/
app.Use(middleware.BearerAuth())

/*
Public user endpoints
*/
account := app.Group("/account")
account.Get("/profile", handler.GetProfile)
// ...

/*
Client endpoints (role: admin)
*/
client := app.Group("/client", middleware.Role("admin"))
client.Get("/", handler.ListClient)
client.Get("/:id", handler.GetClient)

/*
Role endpoints (permission: handle role)
*/
role := app.Group("/role", middleware.Permission("handle role"))
role.Get("/", handler.ListRole)
role.Get("/:id", handler.GetRole)
role.Post("/store", handler.CreateRole)
role.Put("/update/:id", handler.UpdateRole)
role.Delete("/delete/:id", handler.DeleteRole)

```

### 2 Access Token
- JWT
```bash
eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImthbHZlaXJAZ21haWwuY29tIiwicm9sZXMiOlt7ImlkIjoxLCJuYW1lIjoiYWRtaW4iLCJwZXJtaXNzaW9ucyI6W3siaWQiOjEsIm5hbWUiOiJoYW5kbGUgcGVybWlzc2lvbiJ9LHsiaWQiOjIsIm5hbWUiOiJoYW5kbGUgcm9sZSJ9XX1dLCJleHAiOjE3NTQ3OTg0ODEsImlhdCI6MTc1NDc5ODQ0MX0.cmMN7O1_CAEafMjlfgytkn5cr5dcQAnFSwv3nz3r7jbFMpJHvk6RmKYSZFSRvbr1ZqX-rpVRnr6tK9NnMW-aVm-Hzr2Gxrtq8LS_grVLAJRPkxjg7CxsHxA4pta_W7uvuiN3MrWhi1Fd6vHlYJw3pOD0qTvyxXZwyItmKH2DlosKylTPUYXgfQ7pOn5X_Zr7m1RN-M6PYcGgwDaYhQGmgTgfRMM0J3VCS9eSIO5r0wTAl4RjqvrvR5XhMajscF7Ysq6CwymGUDt03e4Uj62BuUCVAc-T6d8Mjz8io59Acq2mBX-CyEIS7-dip22expTZ5hBz7EH1zwhlaUEAvSGhY6PzoDroxA2Gm5uoEcAAzXZMP7g6uim9IiHd6GSeioh7JikXcU4ko-jOSWrDtQgUyZEPMr-_jeGeja8lOdzhh2ZGf88C2P68IpkqKlad53bl1v4FtqxM-JPgYGVa_nVHsBSwLfR9QG0OgOI32etDfTj1jKXc2PCy7qTW2AXTBBsM5jGsIZZBLiTmzBRD7Qa5b7GJbBxwA9qYLF_o_SZa1E5eceXZaxnz2rT5d7kfPHb_gTKR3SIWL1DE9cJ_U5zcdHTiwwBUJFX_czuUSL_UH_b1fX_m7M5L8Yc1A5ri0XKuogxa0ukhj4ASo3hKGoJ24Gg6MrOBcl6KzDLiyGanNz0
```
- Payload
```json
{
  "user_id": 1,
  "email": "nuvantim@gmail.com",
  "roles": [
    {
      "id": 1,
      "name": "admin",
      "permissions": [
        {
          "id": 1,
          "name": "handle permission"
        },
        {
          "id": 2,
          "name": "handle role"
        }
      ]
    }
  ],
  "exp": 1754798481,
  "iat": 1754798441
}

```
## Installation :cd:

### 1. Clone Repository
```bash
git clone https://github.com/Nuvantim/Authrix.git
cd Authrix
```

### 2. Configure Environment
```bash
# Copy environment template
cp .env.local .env

# Edit configuration
nano .env
```

### 3. Environment Variables
```
APP_NAME=Authrix

DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=
DB_NAME=autrix
DB_PORT=5432

MAIL_MAILER=smtp.xxx.com
MAIL_PORT=25
MAIL_USERNAME=user@mail.com
MAIL_PASSWORD=password123
MAIL_FROM_ADDRESS="person@mail.com"

PORT=8080
```

### 4. Install Dependencies
```bash
go mod tidy

```

## Running the Application ⚙️

### Development Mode
```bash
go run cmd/main.go
```

## Deployment 📦
For the deployment process using [Docker](https://www.docker.com/), make sure docker is installed on your server
### 1. Environment Variables
The database configuration is adjusted in the file [docker-compose.yml](https://github.com/Nuvantim/GoStoreAPI/blob/main/docker-compose.yml)
```bash
mv .env.prod .env
```
## 2. Compile project
```bash
make build
```
## 3. Run Docker Compose
```bash
docker compose up -d
```
## 4. Check Logs App
```bash
docker logs authrix_app
```
