# Cafeteria Access Control

A Go-based service for managing and logging access to cafeteria meals using RFID tags. This system supports handling students, meals, cafeterias, and devices with a flexible database layer (MySQL or PostgreSQL).

## Features

- **Access Control**: Verify student access to meals based on RFID tags and schedules.
- **Admin API**: Manage cafeterias, batches, students, meals, and devices.
- **Database Agnostic**: Supports both MySQL and PostgreSQL via configuration.
- **Automatic Migrations**: Database schema is automatically applied on startup.
- **REST API**: Built with `go-chi` for robust routing and middleware support.
- **CORS Support**: Configured for flexible frontend integration.

## Prerequisites

- **Go**: Version 1.24 or higher.
- **Database**: MySQL or PostgreSQL instance.

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/abelmalu/CafeteriaAccessControl.git
   cd CafeteriaAccessControl
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Configure the environment:**
   Create a `.env` file in the root directory. You can use the following template:

   ```env
   # Server Configuration
   SERVER_PORT=8080
   
   # Database Configuration
   DB_TYPE=mysql          # or postgres
   DB_HOST=127.0.0.1
   DB_PORT=3306           # 5432 for postgres
   DB_USER=your_db_user
   DB_Password=your_db_password
   DB_NAME=cafeteria_db

   # File Uploads
   UPLOAD_DIR=./uploads
   ```

   > **Note:** Ensure the `UPLOAD_DIR` exists or the application has permissions to create/access it.

## Usage

1. **Run the application:**
   ```bash
   go run cmd/api/main.go
   ```

   The server will start on the port specified in `SERVER_PORT` (default 8080).
   Database migrations defined in `internal/app/sql/ddl.sql` will be executed automatically upon successful connection.

2. **API Endpoints:**

   **Meal Access:**
   - `GET /api/mealaccess/{sutdentRfid}/{cafeteriaId}` - Attempt meal access.
   - `GET /api/cafeterias` - List all cafeterias.
   - `GET /api/device/verify/{SerialNumber}` - Verify a device.

   **Admin:**
   - `POST /api/admin/create/cafeteria` - Create a new cafeteria.
   - `POST /api/admin/create/batch` - Create a new student batch.
   - `POST /api/admin/create/student` - Register a new student.
   - `POST /api/admin/create/meal` - Define a new meal.
   - `POST /api/admin/register/device` - Register a scanning device.

   **Static Files:**
   - `/static/*` - Serve embedded static assets.
   - `/uploads/*` - Serve uploaded files.

## Project Structure

- `cmd/api`: Application entry point.
- `config`: Configuration loading and validation.
- `internal/api`: HTTP handlers for API endpoints.
- `internal/app`: Core application setup (DB, Router, Wiring).
- `internal/core`: Domain interfaces and business logic definitions.
- `internal/models`: Data structures.
- `internal/repository`: Database implementations (MySQL/Postgres).
- `internal/service`: Business logic implementation.
