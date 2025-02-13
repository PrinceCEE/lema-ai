# Lema Project

## Overview

Lema is a full-stack application built with Golang for the backend and React for the frontend. The backend uses SQLite as the database.

## Prerequisites

- Go (version 1.16 or higher)
- Node.js (version 14 or higher)
- npm (version 6 or higher)

## Setup Instructions

### Backend Setup

1. Navigate to the `api` directory:

   ```sh
   cd api
   ```

2. Copy the `.env.example` file to `.env`:

   ```sh
   cp .env.example .env
   ```

3. Install dependencies:

   ```sh
   go mod tidy
   ```

4. Run the backend server:
   ```sh
   go run ./cmd/api
   ```

### Frontend Setup

1. Navigate to the `web` directory:

   ```sh
   cd web
   ```

2. Install dependencies:

   ```sh
   npm install
   ```

3. Run the frontend development server:
   ```sh
   npm run dev
   ```

## API Endpoints

```
GET    /api/v1/users?limit=x&page=y          // Get users
GET    /api/v1/users/:user_id                // Get a user
GET    /api/v1/users/count                   // Get user's count
POST   /api/v1/posts                         // Create a post
GET    /api/v1/posts?user_id=x               // Get a user's posts
POST   /api/v1/posts/:post_id                // Get a post
DELETE /api/v1/posts/:post_id                // Delete a post
```

## Running the Project Locally

1. Open two terminal windows or tabs.

2. In the first terminal, navigate to the `api` directory and start the backend server:

   ```sh
   cd api
   go run ./cmd/api
   ```

````

3. In the second terminal, navigate to the `web` directory and start the frontend development server:

   ```sh
   cd web
   npm run dev
   ```

4. Open your browser and go to `http://localhost:3000` to view the application.

## Unit Tests

### Backend Tests

1. Navigate to the `api` directory:

   ```sh
   cd api
   ```

2. Run the tests:
   ```sh
   go test ./...
   ```

### Frontend Tests

1. Navigate to the `web` directory:

   ```sh
   cd web
   ```

2. Run the tests:
   ```sh
   npm test
   ```

## Directory Structure

```
lema/
├── api/                # Backend source code
│   ├── cmd/
│   │   └── api/        # Main entry point for the backend server
│   ├── internal/       # Internal packages for the backend
│   ├── models/         # Database models
│   ├── handlers/       # HTTP handlers
│   └── ...             # Other backend files
├── web/                # Frontend source code
│   ├── src/
│   │   ├── components/ # React components
│   │   ├── pages/      # Next.js pages
│   │   ├── api/        # API service files
│   │   └── ...         # Other frontend files
├── README.md           # This file
└── Makefile            # Makefile for running the backend and frontend
```

## Makefile

You can also use the Makefile to run the backend and frontend servers:

1. Run the backend server:

   ```sh
   make api
   ```

2. Run the frontend development server:
   ```sh
   make web
   ```
````
