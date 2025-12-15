# User Management API

A RESTful API built using Go and Fiber to manage users with `name` and `dob`. User age is calculated dynamically when fetching user data and is not stored in the database.


## Tech Stack

- Go
- Fiber
- PostgreSQL (Supabase)
- SQLC
- go-playground/validator
- Uber Zap
- Docker


## Prerequisites

- Go 1.24+
- PostgreSQL or Supabase project
- Docker (optional)
- sqlc installed


## Setup and Run 

1. Clone the repository and move into the project directory.

```
git clone https://github.com/SudheshHegde2002/user-api.git
cd user-api
```

2. Create a `.env` file in the project root and add the database connection string.

```env
DATABASE_URL=postgresql://postgres.mdebyejtogorxwvcdnls:user-api-test1234@aws-1-ap-southeast-1.pooler.supabase.com:6543/postgres?pgbouncer=true
```

3. Apply database migrations by running the SQL files present in the `db/migrations` directory on your PostgreSQL database.

4. Generate SQLC code.

```bash
sqlc generate
```

5. Start the server locally.

```bash
go run cmd/server/main.go
```

The API will be available at `http://localhost:3000`.


## Run with Docker

6. Build the Docker image.

```bash
docker build -t user-api .
```

7. Run the Docker container.

```bash
docker run -p 3000:3000 \
  -e DATABASE_URL="postgresql://postgres.mdebyejtogorxwvcdnls:user-api-test1234@aws-1-ap-southeast-1.pooler.supabase.com:6543/postgres?pgbouncer=true" \
  user-api
```


## API Endpoints

- POST /users
- GET /users/:id
- GET /users?limit=10&offset=0
- PUT /users/:id
- DELETE /users/:id

