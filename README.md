# GEFZI

## Project Description
GEFZI is a student project web application that allows multiple users to find common free time slots. The focus is on identifying the next available time for scheduling events rather than the next scheduled appointment.

## Repository
[https://github.com/emilygdu/gefzi](https://github.com/emilygdu/gefzi)

## Requirements
- Go 1.25.2  
- SQLite database file excluded in the repository 
- No additional dependencies required (Node.js/npm not needed)

## Structure

```bash
├── cmd/
│   └── server/
│       └── main.go
│
├── db/
│   └── gefzi.db
│
├── internal/
│   ├── database/
│   │   └── connection.go
│   │
│   ├── models/
│   │   ├── event.go
│   │   ├── group_calendar.go
│   │   └── user.go
│   │
│   └── routes/
│       ├── availability_helpers.go
│       ├── availability_routes.go
│       ├── event_routes.go
│       ├── group_calendar_routes.go
│       └── user_routes.go
│
├── go.mod
├── go.sum
└── README.md
```

## Installation
1. Clone the repository:  
   ```
   git clone https://github.com/emilygdu/gefzi
   cd gefzi
   ```

2. Install Go dependencies:
    ```
    go mod tidy
    ```

3. Start the server (the included SQLite database will be used automatically):
    ```
    go run main.go
    ```

The server runs on http://localhost:8080

## API Endpoints
#### Users
GET /users → List all users

POST /users → Create a new user

Example JSON Body:
```json
{
  "first_name": "Alice",
  "last_name": "Example",
  "email": "alice@example.com",
  "group_calendar_id": 1
}
```

#### Group Calendars
GET /groupcalendars → List all group calendars

POST /groupcalendars → Create a new group calendar

Example JSON Body:
```json
{
  "name": "Project Team",
  "work_start": "08:00",
  "work_end": "17:00",
  "weekend_blocked": true
}
```

#### Events
Events represent blocked time slots in a group calendar.
They are used to calculate free time (availability) and are not returned as free slots.

GET /events → List all events

POST /events → Create a new event

Example JSON Body:
```json
{
  "date": "2026-01-29",
  "start_time": "09:00",
  "end_time": "10:00",
  "visibility": "private",
  "group_calendar_id": 1
}
```

#### Availability
GET /availability/month

Query Parameters:
- group_calendar_id (int)
- year (int)
- month (int, 1–12)

Example:
GET /availability/month?group_calendar_id=1&year=2026&month=1

Response:
Returns all free time slots per day within the defined working hours.



## Notes
- The SQLite database file is not tracked in the repository.
- A new database is created automatically on server start.
- Database tables are auto-migrated on startup using GORM.
- Events represent blocked time slots and are used to calculate availability.
- Both private and business events block time equally; the distinction is only relevant for frontend visualization.
- CORS is enabled to allow frontend development and testing.


