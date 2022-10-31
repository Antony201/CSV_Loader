# Csv Loader App

---
The api for csv loader (uploading csv to save to db and filters to search file data in db)

## Dependencies:

---
- go >= 1.16
- docker >= 19

## Build and run

---
1. Create .env file in project folder and set variables like in .default.env (DB_PASSWORD, DATABASE_URL)
2. Build project
> make build

2. If you are running for the first time:
> make migrate

3. Run Project
> make run

---

## Api Description

Api is deskcribed in swagger docs by url 
> hostname:8000/swagger/index.html

---
## Testing

To run test use:
> make test

---
## Migrations
To up migrations use:
> make migrate

To dawn:
> make migrate-down

---

### Goals:

- Increase test covering to 80-100%
- Optimize project structure

---

#### Author: Anton Nebozhynskyi