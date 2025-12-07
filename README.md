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
    curl -X GET 'http://localhost:8080/users'

### Get single User
    curl -X GET 'http://localhost:8080/users?id=7'

### Create
    curl -X POST 'http://localhost:8080/users' \
    -d '{"name": "UsuarioPrueba", "email": "usuarioprueba@gmail.com", "password": "password1", "rol_id": 1}'

### Update
    curl -X PUT 'http://localhost:8080/users' \
    -d '{"id": 7, "name": "userUpdated", "email": "useremailupdated@gmail.com", "password": "passwordupdated", "rol_id": 2}'

### Delete
    curl -X DELETE 'http://localhost:8080/users' \
    -d '{"id": 7}'
