# GEFZI

## Project Description
GEFZI is a student project web application that allows multiple users to find common free time slots. The focus is on identifying the next available time for scheduling events rather than the next scheduled appointment.

## Repository
[https://github.com/emilygdu/gefzi](https://github.com/emilygdu/gefzi)

## Requirements
- Go 1.25.2  
- SQLite database file included in the repository (`db.sqlite`)  
- No additional dependencies required (Node.js/npm not needed)

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
`GET /users` → List all users (including their group calendars)
`POST /users` → Create a new user

Example JSON Body for POST /users:

```
{
  "name": "Alice",
  "email": "alice@example.com",
  "group_calendar_id": 1
}
```

#### Events
`GET /events` → List all events
`POST /events` → Create a new event

Example JSON Body for POST /events:

```
{
  "title": "Team Meeting",
  "start_time": "2025-11-22T10:00:00",
  "end_time": "2025-11-22T11:00:00",
  "user_id": 1
}
```

#### GroupCalendars
`GET /groupcalendars` → List all group calendars (including members)
`POST /groupcalendars` → Create a new group calendar

Example JSON Body for POST /groupcalendars:

```
{
  "name": "Projectteam"
}
```

## Notes
The included SQLite database file (db.sqlite) already contains sample data. No additional setup is required.

The server automatically migrates the tables for User and Event on start.

CORS is enabled for frontend testing.

