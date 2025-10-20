# companies

Manage CRU for a single table company with Go Lang and Gin framework.

## Setup

To run this project locally, ensure you have Go  (1.25) installed and run:

```bash
go run main.go
```

## Docker

To build and run the Docker container for production, use the following commands:

```bash
docker build -f Dockerfile.production -t companies-app .
docker run -p 8089:8080 companies-app
```

Then hit `http://localhost:8089/ping` in your browser.

## Run tests

To run test, use the following command:

```bash
go test ./...
```
### Test Coverage

To check test coverage, use the following command:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
This will generate a coverage report and open it in your default web browser. Currently, the test coverage is negligible.

## Database, table and sample data

This project uses MySQL as the database. You can get a free working MySQL database at [Aiven](https://aiven.io/free-mysql-database).
To create the `company` table and insert sample data, use the following SQL commands:

```sql
mysql -u your_username -p -h your_host -P your_port your_database_name
source data/init.sql;
```

Or you can use [Docker and Docker compose for a MySQL](https://geshan.com.np/blog/2022/02/mysql-docker-compose/) database.

## Environment Variables

The example environment variables are in the `.env.example` file. Create a `.env` file in the root directory and set the right environment variables.
For instance if you want to run the app locally in a different port (than 8080), you can set the `PORT` variable in the `.env` file:

```env
PORT=8089
```
