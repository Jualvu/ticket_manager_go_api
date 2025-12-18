### Backend API for a simple ticket manager software

## Setup

1. Install Go

2. Create the sqlite database:
    2.1. Install sqlite3
    * For macs only
    2.2. Create the db file: touch ticket_system.db
    2.3. Create the db schema: sqlite3 ticket_system.db < schema.sql
    2.4. Optional** install database seed sample: sqlite3 ticket_system.db < seed.sql

    * For windows
     2.2

     
3. Ru the project depending on use case: 
    - for dev -> go run cmd/api/main.go
    - for prod -> 
        1. construct binary with: go build cmd/api/main.go
        2. Run the binary: ./main



### Curl requests for testing

## Users
**NOTE: change id for a valid value

### Get All
    curl -X GET 'http://localhost:8080/users/'

### Get single User
    curl -X GET 'http://localhost:8080/users/3'

### Create
    curl -X POST 'http://localhost:8080/users/' \
    -d '{"name": "UsuarioPrueba", "email": "usuarioprueba@gmail.com", "password": "password1", "rol_id": 1}'

### Update
    curl -X PUT 'http://localhost:8080/users/4' \
    -d '{"name": "userUpdated", "email": "useremailupdated@gmail.com", "password": "passwordupdated", "rol_id": 2}'

### Delete
    curl -X DELETE 'http://localhost:8080/users/4'





## Tickets
**NOTE: change id for a valid value

### Get All
    curl -X GET 'http://localhost:8080/tickets/'

### Get single User
    curl -X GET 'http://localhost:8080/tickets/3'

### Create
    curl -X POST 'http://localhost:8080/tickets/' \
    -d '{"title": "New ticket for testing create", "description": "This is a new ticket description", "state_id": 1, "priority_id": 1, "assigned_to_user_id": 1, "created_by_user_id": 2}'

### Update
    curl -X PUT 'http://localhost:8080/tickets/4' \
    -d '{"title": "New ticket testing Update", "description": "This is a ticket to test update", "state_id": 2, "priority_id": 2, "assigned_to_user_id": 3}'

### Delete
    curl -X DELETE 'http://localhost:8080/tickets/4'




## Comments
**NOTE: change id for a valid value

### Get All
    curl -X GET 'http://localhost:8080/comments/'

### Get single User
    curl -X GET 'http://localhost:8080/comments/3'

### Create
    curl -X POST 'http://localhost:8080/comments/' \
    -d '{"content": "New comment for testing create", "created_by_user_id": 2, "ticket_id": 1}'

### Update
    curl -X PUT 'http://localhost:8080/comments/4' \
    -d '{"content": "New comment testing update"}'

### Delete
    curl -X DELETE 'http://localhost:8080/comments/4'



## Auth

### Login 

# normal user
    curl -X POST 'http://localhost:8080/auth/login' \
    -d '{"email": "bob@example.com", "password": "hashed_password_1"}'

# admin
    curl -X POST 'http://localhost:8080/auth/login' \
    -d '{"email": "alice@example.com", "password": "hashed_password_1"}'