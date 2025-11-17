### Backend API for a simple ticket manager software

## Setup

1. Install Go

2. Create the sqlite database:
    2.1. Install sqlite3
    2.2. Create the db file: touch ticket_system.db
    2.3. Create the db schema: sqlite3 ticket_system.db < schema.sql
    2.4. Optional** install database seed sample: sqlite3 ticket_system.db < seed.sql

3. Run the project depending on use case: 
    - for dev -> go run cmd/api/main.go
    - for prod -> 
        1. construct binary with: go build cmd/api/main.go
        2. Run the binary: ./main

